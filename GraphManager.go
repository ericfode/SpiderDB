package spiderDB

//TO-DO:
//   
//
//  + make type keys and type set
//  + make sorted set of all nodes with type id as score
//  + when inserting node, add to sorted set 
//  + add register/deregister type functions
//      + deregister should only work if no nodes of that type exist
//      + register should only work if type does not already exist

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
	//	nodeConstructors []NodeConstructor
	//	edgeConstructors []EdgeConstructor
}

func (gm *GraphManager) Initialize() {
	//var e error

	gm.client, _ = redis.NewSynchClient()
	gm.client.Set(currIndex_s, []byte("0"))

	//gm.Connect(port, ipAddr, dbNum)

	gm.nodes = make(map[string]Node)
	gm.edges = make(map[string]Edge)
}

/*
func (gm *GraphManager) Connect(port int, ipAddr string, dbNum int) {
	gm.client, _ = redis.NewSynchClient()
}
*/
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
	n.SetID(nindex)
	gm.nodes[nindex] = n
}

func (gm *GraphManager) DeleteNode(n Node) {
	//Database
	nindex := n.GetID()
	gm.client.Srem(nodes_s, []byte(nindex))
	gm.client.Del(nindex)
	//Local
	gm.nodes[nindex] = nil
	n = nil
}

func (gm *GraphManager) FindNode(nindex string) Node {
	n, ok := gm.nodes[nindex]
	//if local
	if ok == true {
		return n
	}
	//otherwise, get from db
	return gm.GetNode(nindex)
}

//bulk update - pushes everything to database
func (gm *GraphManager) UpdateNode(n Node) error {
	n.GetID()

	//e := gm.client.Hset(n.GetID(), props_s, n.GetPropMap())
	return nil
}

//pushes change of single prop to db
func (gm *GraphManager) UpdateNodeProp(n Node, prop string, value []byte) error {

	if n.GetId() == "" {
		return &NodeNotAddedToDBError{e}
	}

	nindex := node_s + n.GetId()

	gm.client.Hset(nindex, prop, value)
}

func (gm *GraphManager) GetNode(index string) Node {
	nodeIdx, err := gm.client.Hget(node_s+index, props_s)

	if err != nil || nodeIdx == nil {
		return nil
	}

	node := new(Node)
	node.SetID(nodeIdx)
	gm.nodes[nodeIdx] = node
	return node
}

//func (gm *GraphManager) GetAdjPairs(node *Node) *[]AdjPair {}

//Add neigbhbor both locally and in db

//attach bidirectional Neighbor
//think about using this same pattern (passing a type) for other funcs 
func (gm *GraphManager) Attach(node1 Node, node2 Node, edge Edge) {
	//decide if being able to add nodes and edges that don't have ids
	//into the data base is a good idea or if making the user explicitly
	//do it is a good idea
	//i think that having different behavior then just attaching a neightbor
	//is a bad idea
	gm.client.Hset(node_s+node1.GetID()+adj_s, node2.GetID(), e.GetId())
	if !e.IsDirected() {
		gm.client.Hset(node_s+node2.GetID()+adj_s, node1.GetID(), e.GetId())
	}
	node1.AddEdges(edge)
	node2.AddEdges(edge)
	edge.SetFirstNode(node1)
	edge.SetSecondNode(node2)
}

func (gm *GraphManager) GetNeighbors(node Node) ([]*Connection, error) {
	conns, ok := gm.client.Hgetall(key)
	if ok != nil {
		return _, ok.Error()
	}
	for index, val := range conns {

	}

}

//EDGE MANAGEMENT

func (gm *GraphManager) AddEdge(e Edge) {

	//Database
	index := gm.GetNextIndex()
	nindex := edge_s + index
	props := e.GetPropMap()

	//Add edge props to database edge
	for k, v := range props {
		gm.client.Hset(eindex, k, v)
	}

	//Add node to index
	gm.client.Sadd(edges_s, []byte(eindex))

	//Local
	e.SetId(eindex)
	gm.edges[eindex] = e
}

func (gm *GraphManager) DeleteEdge(e Edge) {

	eindex = e.GetId()

	//remove locally
	gm.edges[eindex] = nil

	//remove from database
	gm.client.Del(eindex)

	//remove from database's edge-index
	gm.client.Srem(edges_s, eindex)

	e = nil
}

func (gm *GraphManager) FindEdge(id int) Edge {
	return nil
}

func (gm *GraphManager) UpdateEdge(e Edge) bool {
	return true
}

func (gm *GraphManager) UpdateEdgeProp(e Edge, prop string, value []byte) error {
	if e.GetId() == nil {
		return &EdgeNotAddedToDBError{e}
	}

	nindex := edge_s + e.GetId()

	gm.client.Hset(nindex, prop, value)
}

func (gm *GraphManager) GetEdge(id int) Edge {
	return nil
}

func (gm *GraphManager) GetNodeEdges(n Node) map[string][]Edge {

	ret := make(map[string][]*Edge, len(gm.edges))
	//for each edge, classify and add edge pointer to correct slice
	for i := 0; i < len(gm.edges); i++ {
		typ := gm.edges[i].GetType()
		ret[typ] = append(ret[typ], gm.edges[i])
	}
	return ret
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
