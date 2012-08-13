package spiderDB

//TO-DO:
//   
//  + make type keys and type set
//  + make sorted set of all nodes with type id as score
//  + when inserting node, add to sorted set 

import "github.com/alphazero/Go-Redis"

// consts (typo prevention)
const currIndex_s = "currIndex"
const node_s = "node:"
const edge_s = "edge:"
const nodes_s = "nodes"
const edges_s = "edges"
const props_s = "props"
const adj_s = "adj"

//Initializes database at startup

type GraphManager struct {
	nodes  map[string]Node
	edges  map[string]Edge
	client redis.Client
}

func (gm *GraphManager) Initialize() {
	//var e error

	gm.client, _ = redis.NewSynchClient()
	gm.client.Set(currIndex_s, []byte("0"))

	gm.Connect(6379, "127.0.0.1", 13)

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
	nindex := node_s+n.GetID()
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

func (gm *GraphManager) NodeFromHash(hash [][]byte, construct NodeConstructor) (Node, bool){
	node := construct(string(hash[1]),gm)
	propMap := make(map[string][]byte)
	for i := 0; i < len(hash); i +=2{
		propMap[string(hash[i])] = hash[i+1]
	}
	node.SetPropMap(propMap)
	return node, true
}

func (gm *GraphManager) GetNode(index string, construct NodeConstructor) (Node, error) {
	if exists , err := gm.client.Sismember(nodes_s,[]byte(node_s + index)); exists != true{
		return nil, &KeyNotFoundError{index}
	} else if err != nil{
		return nil, err
	}

	hash, err := gm.client.Hgetall(node_s+index)
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
	gm.nodes[node_s+node.GetID()] = node
	return node, nil

}

//func (gm *GraphManager) GetAdjPairs(node *Node) *[]AdjPair {}

//Add neigbhbor both locally and in db

//attach bidirectional Neighbor
//think about using this same pattern (passing a type) for other funcs 
func (gm *GraphManager) Attach(node1 Node, node2 Node, e Edge) {
	//decide if being able to add nodes and edges that don't have ids
	//into the data base is a good idea or if making the user explicitly
	//do it is a good idea
	//i think that having different behavior then just attaching a neightbor
	//is a bad idea
	gm.client.Hset(node_s+node1.GetID()+adj_s, node2.GetID(), []byte(e.GetID()))
	if !e.IsDirected() {
		gm.client.Hset(node_s+node2.GetID()+adj_s, node1.GetID(), []byte(e.GetID()))
	}
	node1.AddEdges([]Edge{e})
	node2.AddEdges([]Edge{e})
	e.SetFirstNode(node1)
	e.SetSecondNode(node2)

}

func (gm *GraphManager) GetNeighbors(node Node) []Connection {
	//	conns, _ := gm.client.Hgetall(node_s + node.GetID())

	//	for _, _ := range conns {

	//	}
	return nil
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

func (gm *GraphManager) FindEdge(id int) Edge {
	return nil
}

func (gm *GraphManager) UpdateEdge(e Edge) bool {
	return true
}

func (gm *GraphManager) UpdateEdgeProp(e Edge, prop string, value []byte) error {
	if e.GetID() == "" {
		//		return &EdgeNotAddedToDBError{e}
	}

	nindex := edge_s + e.GetID()

	gm.client.Hset(nindex, prop, value)
	return nil
}

func (gm *GraphManager) GetEdge(id int) Edge {
	return nil
}

func (gm *GraphManager) GetNodeEdges(n Node) map[string][]Edge {

	//	ret := make(map[string][]Edge, len(gm.edges))
	//for each edge, classify and add edge pointer to correct slice
	/*
		for i := 0; i < len(gm.edges); i++ {
			typ := gm.edges[i].GetType()
			ret[typ] = append(ret[typ], gm.edges[i])
		}
		return ret
	*/
	return nil
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
