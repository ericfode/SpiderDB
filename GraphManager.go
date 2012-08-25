package spiderDB

//TO-DO:
//   
//  + make type keys and type set
//  + make sorted set of all nodes with type id as score
//  + when inserting node, add to sorted set 

import "github.com/alphazero/Go-Redis"
import "strings"

// consts (typo prevention)
const currIndex_s = "currIndex"
const node_s = "node:"
const edge_s = "edge:"
const nodes_s = "nodes"
const edges_s = "edges"
const props_s = "props"
const adj_s = "adj"
const inadj_s = "inadj"

//Initializes database at startup

type GraphManager struct {
	nodes  map[string]Node
	edges  map[string]Edge
	client redis.Client
}

func (gm *GraphManager) Initialize() {
	//var e error

	gm.Connect(6379, "127.0.0.1", 13)
	gm.client.Set(currIndex_s, []byte("0"))

	gm.nodes = make(map[string]Node)
	gm.edges = make(map[string]Edge)
}

func (gm *GraphManager) Connect(port int, ipAddr string, dbNum int) {
	gm.client, _ = redis.NewSynchClient()
}

func (gm *GraphManager) GetNextIndex() string {
	index, _ := gm.client.Get(currIndex_s)
	gm.client.Incr(currIndex_s)

	return string(index)
}

func (gm *GraphManager) GetCurIndex() string {
	index, _ := gm.client.Get(currIndex_s)
	return string(index)
}

// NODE MANAGEMENT

func (gm *GraphManager) AddNode(n Node) {
	//Database

	index := gm.GetNextIndex()
	nindex := node_s + index
	props := n.GetPropMap()

	//Add node props to database node
	for k, v := range props {
		gm.client.Hset(nindex, k, v)
	}

	//Add node to node-index
	gm.client.Sadd(nodes_s, []byte(nindex))

	//Local
	n.SetID(index)
	gm.nodes[nindex] = n
}

func (gm *GraphManager) DeleteNode(n Node) {
	//Database
	nindex := node_s + n.GetID()
	gm.client.Srem(nodes_s, []byte(nindex))
	gm.client.Del(nindex)
	//Local
	gm.nodes[nindex] = nil
	n = nil
}

func (gm *GraphManager) FindNode(nindex string, construct NodeConstructor) (Node, error) {
	n, ok := gm.nodes[nindex]
	//if local
	if ok == true {
		return n, nil
	}
	//otherwise, get from db

	return gm.GetNode(nindex, construct)
}

func (gm *GraphManager) FindNodeWithValue(hashName string, value string, construct NodeConstructor) ([]Node, error) {
	nodeIDs, _ := gm.client.Smembers(nodes_s)
	nodes := make([]Node, 0)
	var err error
	for _, v := range nodeIDs {
		var newNode Node
		idx := strings.LastIndex(string(v), ":")

		val, _ := gm.client.Hget(string(v), hashName)
		valstr := string(val)
		print(string(v) + " :")
		println(valstr)
		if value == valstr {
			newNode, err = gm.GetNode(string(v)[idx+1:], construct)
			nodes = append(nodes, newNode)
			if err != nil {
				return nil, err
			}
		}

	}
	return nodes, nil
}

//bulk update - pushes everything to database
func (gm *GraphManager) UpdateNode(n Node) error {

	for k, v := range n.GetPropMap() {
		if e := gm.UpdateNodeProp(n, k, v); e == nil {
			return e
		}
	}
	return nil
}

//pushes change of single prop to db
func (gm *GraphManager) UpdateNodeProp(n Node, prop string, value []byte) error {
	id := n.GetID()
	if id == "" {
		return &dbError{"Not added to DB"}
	}

	nindex := node_s + id

	gm.client.Hset(nindex, prop, value)
	return nil
}

//TODO : breakout node_s + index 
//TODO : rename nodes_s to somthing less ambiguous

func (gm *GraphManager) NodeFromHash(hash [][]byte, construct NodeConstructor) (Node, bool) {
	node := construct(string(hash[1]), gm)
	propMap := make(map[string][]byte)
	for i := 0; i < len(hash); i += 2 {
		propMap[string(hash[i])] = hash[i+1]
	}
	node.SetPropMap(propMap)
	return node, true
}

func (gm *GraphManager) GetNode(index string, construct NodeConstructor) (Node, error) {
	if exists, err := gm.client.Sismember(nodes_s, []byte(node_s+index)); exists != true {
		return nil, &KeyNotFoundError{index}
	} else if err != nil {
		return nil, err
	}

	hash, err := gm.client.Hgetall(node_s + index)
	if err != nil {
		return nil, err
	} else if hash == nil {
		return nil, &KeyNotFoundError{index}
	}

	// Need to add constructor stuff
	node, ok := gm.NodeFromHash(hash, construct)
	if !ok {
		return nil, &dbError{"Something that should not be able to brake broke???"}
	}
	node.SetID(index)
	gm.nodes[node_s+node.GetID()] = node
	return node, nil

}

//func (gm *GraphManager) GetAdjPairs(node *Node) *[]AdjPair {}

func (gm *GraphManager) Attach(node2 Node, node1 Node, e Edge) {
	gm.client.Hset(node_s+node1.GetID()+adj_s, node2.GetID(), []byte(e.GetID()))
	gm.client.Hset(node_s+node2.GetID()+inadj_s, node1.GetID(), []byte(e.GetID()))
	if !e.IsDirected() {
		gm.client.Hset(node_s+node2.GetID()+adj_s, node1.GetID(), []byte(e.GetID()))
		gm.client.Hset(node_s+node1.GetID()+inadj_s, node2.GetID(), []byte(e.GetID()))
	}

	node1.AddEdge(e)
	node2.AddEdge(e)
	e.SetFirstNode(node1)
	e.SetSecondNode(node2)

}

func (gm *GraphManager) GetNeighbors(node Node, constE EdgeConstructor, constN NodeConstructor) ([]Connection, error) {

	adjId := node_s + node.GetID() + adj_s
	inadjId := node_s + node.GetID() + inadj_s

	adjArray, _ := gm.client.Hgetall(adjId)
	inadjArray, _ := gm.client.Hgetall(inadjId)

	neighbors := make([]Connection, 0)
	if adjArray != nil {
		for k, v := range ByteAAtoStringMap(adjArray) {
			nb, err := gm.GetNode(k, constN)
			if err != nil {
				return nil, err
			}
			ec, err := gm.GetEdge(string(v), constE)
			if err != nil {
				return nil, err
			}

			newConn := Connection{NodeA: node, NodeB: nb, Edg: ec}
			neighbors = append(neighbors, newConn)
		}
	}
	if inadjArray != nil {
		for k, v := range ByteAAtoStringMap(inadjArray) {
			nb, err := gm.GetNode(k, constN)
			if err != nil {
				return nil, err
			}
			ec, err := gm.GetEdge(string(v), constE)
			if err != nil {
				return nil, err
			}

			newConn := Connection{NodeA: nb, NodeB: node, Edg: ec}
			neighbors = append(neighbors, newConn)
		}
	}
	return neighbors, nil
}

//EDGE MANAGEMENT

func (gm *GraphManager) AddEdge(e Edge) {

	//Database
	index := gm.GetNextIndex()
	eindex := edge_s + index
	props := e.GetPropMap()

	//Add edge props to database edge
	for k, v := range props {
		gm.client.Hset(eindex, k, v)
	}

	//Add node to index
	gm.client.Sadd(edges_s, []byte(eindex))

	//Local
	e.SetID(string(index))
	gm.edges[eindex] = e
}

func (gm *GraphManager) DeleteEdge(e Edge) {

	eindex := e.GetID()

	//remove locally
	gm.edges[eindex] = nil

	//remove from database
	gm.client.Del(eindex)

	//remove from database's edge-index
	gm.client.Srem(edges_s, []byte(eindex))

	e = nil
}

func (gm *GraphManager) FindEdge(id string, construct EdgeConstructor) (Edge, error) {
	eindex := edge_s + id
	//if locally available
	if e, ok := gm.edges[eindex]; ok == true {
		return e, nil
	}

	//else extract from db (GetEdge saves local copy)
	edge, err := gm.GetEdge(eindex, construct)
	return edge, err
}

// pushes changes of e(local edge) to db
func (gm *GraphManager) UpdateEdge(e Edge) error {

	for k, v := range e.GetPropMap() {
		if err := gm.UpdateEdgeProp(e, k, v); err != nil {
			return err
		}
	}
	return nil
}

func (gm *GraphManager) UpdateEdgeProp(e Edge, prop string, value []byte) error {
	if e.GetID() == "" {
		return &dbError{"Edge not added to DB"}
	}

	eindex := edge_s + e.GetID()

	gm.client.Hset(eindex, prop, value)
	return nil
}

func (gm *GraphManager) EdgeFromHash(hash [][]byte, construct EdgeConstructor) (Edge, bool) {
	edge := construct(string(hash[1]), gm)
	propMap := make(map[string][]byte)
	for i := 0; i < len(hash); i += 2 {
		propMap[string(hash[i])] = hash[i+1]
	}
	edge.SetPropMap(propMap)
	return edge, true
}

func (gm *GraphManager) GetEdge(id string, construct EdgeConstructor) (Edge, error) {
	eindex := edge_s + id
	// check if edge is even in database
	exists, err := gm.client.Sismember(edges_s, []byte(eindex))
	if exists != true {
		return nil, &KeyNotFoundError{string(id)}
	} else if err != nil {
		return nil, err
	}

	//retrieve edge's properties, if in db
	hash, err := gm.client.Hgetall(eindex)
	if err != nil {
		return nil, err
	}

	edge, ok := gm.EdgeFromHash(hash, construct)
	if !ok {
		return nil, err
	}

	// save edge locally 
	gm.edges[string(edge.GetID())] = edge
	return edge, nil
}

func (gm *GraphManager) GetNodeEdges(n Node, construct EdgeConstructor, adjId string) []Edge {
	hash, err := gm.client.Hgetall(adjId)

	edges := make([]Edge, len(hash)/2)

	if err != nil {
		return nil
	}

	for i := 0; i < len(hash); i += 2 {
		eindex := edge_s + string(hash[i+1])
		e, _ := gm.GetEdge(eindex, construct)
		edges[i/2] = e
	}

	return edges
}

func (gm *GraphManager) GetOutgoingNodeEdges(n Node, construct EdgeConstructor) []Edge {
	adjId := node_s + n.GetID() + adj_s
	return gm.GetNodeEdges(n, construct, adjId)
}

func (gm *GraphManager) GetIncomingNodeEdges(n Node, construct EdgeConstructor) []Edge {
	inadjId := node_s + n.GetID() + inadj_s
	return gm.GetNodeEdges(n, construct, inadjId)
}

func (gm *GraphManager) GetAllNodeEdges(n Node, construct EdgeConstructor) []Edge {
	adjId := node_s + n.GetID() + adj_s
	inadjId := node_s + n.GetID() + inadj_s
	edgesOut := gm.GetNodeEdges(n, construct, adjId)
	edgesIn := gm.GetNodeEdges(n, construct, inadjId)

	allEdges := append(edgesIn, edgesOut...)
	return allEdges
}

func (gm *GraphManager) ClearAll() {
	gm.client.Set(currIndex_s, []byte("0"))
	gm.client.Flushdb()
	gm.nodes = nil
	gm.edges = nil
	gm.client = nil
}

//PRIVATE FUNCTIONS

//func ( gm *GraphManager ) addAdjEntry( )
//func ( gm *GraphManager ) removeAdjEntry( )

//FIXME: Nodes of multiple types will make this blow up need to do type 
//thing talked about at the top of this file
func (gm *GraphManager) GetAllNodes(construct NodeConstructor) ([]Node, error) {
	nodeIDs, _ := gm.client.Smembers(nodes_s)
	nodes := make([]Node, len(nodeIDs))
	var err error
	for i, v := range nodeIDs {
		idx := strings.LastIndex(string(v), ":")
		nodes[i], err = gm.GetNode(string(v)[idx+1:], construct)
		if err != nil {
			return nil, err
		}
	}
	return nodes, nil
}
