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

//Initializes database at startup

type GraphManager struct {
	nodes  map[string]*Node
	edges  map[string]*Edge
	client redis.Client
}

func (gm *GraphManager) Initialize() {
	//var e error
	gm.client, _ = redis.NewSynchClient()
	gm.client.Set("currIndex", []byte("0"))

	gm.nodes = make(map[string]*Node)
	gm.edges = make(map[string]*Edge)
}

func (gm *GraphManager) Connect(port int, ipAddr string, dbNum int) {
}

func (gm *GraphManager) GetNextIndex() string {
	index, _ := gm.client.Get("currIndex")
	gm.client.Incr("currIndex")

	return string(index)
}

// NODE MANAGEMENT

func (gm *GraphManager) AddNode(n *Node) {
	//Database
	index := gm.GetNextIndex()
	props := n.GetPropMap()

	gm.client.Hset("node:"+index, "props", props[index])
	//gm.client.Hset("node:"+string(index), "index", []byte(index))

	//Add node to index
	gm.client.Sadd("nodes", props[string(index)])
	//gm.client.Sadd("nodes", []byte(index))

	//Local
	n.SetID(index)
	gm.nodes[string(index)] = n
}

func (gm *GraphManager) DeleteNode(n *Node) {
	//Database
	index = n.GetID()
	gm.client.Srem("nodes", []byte(index))
	gm.client.Del("node:" + index)
	//Local
	gm.nodes[index] = nil
	n = nil
}

func (gm *GraphManager) FindNode(index string) *Node {
	n, ok := gm.nodes[index]
	//if local
	if ok == true {
		return n
	}
	//otherwise, get from db
	return gm.GetNode(index)
}

func (gm *GraphManager) UpdateNode(n *Node) bool {
	return true
}

func (gm *GraphManger) UpdateNodeProp(n *Node, prop String, value []byte) {

}

func (gm *GraphManager) GetNode(index string) *Node {
	nodeIdx, err := gm.client.Hget("node:"+index, "props")

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
func (gm *GraphManager) AddNeighbor(node *Node, edge *Edge) {

}

func (gm *GraphManager) GetNeighbors(node *Node) []*Node {

}

//EDGE MANAGEMENT

func (gm *GraphManager) AddEdge(e *Edge) {

	index := gm.GetNextIndex()
	props := e.GetWeight()
	gm.client.Hset("edge:"+index, "props", []byte(props))

	//Add node to index
	gm.client.Sadd("edges", []byte(props))

	//Local
	e.SetId(index)
	gm.edges[index] = e
}

func (gm *GraphManager) DeleteEdge(e *Edge) {
}

func (gm *GraphManager) FindEdge(id int) *Edge {
	return nil
}

func (gm *GraphManager) UpdateEdge(e *Edge) bool {
	return true
}

func (gm *GraphManager) GetEdge(id int) *Edge {
	return nil
}

func (gm *GraphManager) GetNodeEdges(n *Node) map[string][]*Edge {

	ret := make(map[string][]*Edge, len(gm.edges))
	//for each edge, classify and add edge pointer to correct slice
	for i := 0; i < len(gm.edges); i++ {
		typ := gm.edges[i].GetType()
		ret[typ] = append(ret[typ], gm.edges[i])
	}
	return ret
}

func (gm *GraphManager) ClearAll() {
	gm.client.Set("currIndex", []byte("0"))
	gm.client.Flushdb()
	gm.nodes = nil
	gm.edges = nil
	gm.client = nil
}

//PRIVATE FUNCTIONS

//func ( gm *GraphManager ) addAdjEntry( )
//func ( gm *GraphManager ) removeAdjEntry( )
