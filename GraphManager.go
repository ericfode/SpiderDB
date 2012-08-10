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
	gm.client.Set("nodeCount", []byte("0"))
	gm.nodes = make(map[string]*Node)
	gm.edges = make(map[string]*Edge)
}

func (gm *GraphManager) Connect(port int, ipAddr string, dbNum int) {
}

// NODE MANAGEMENT

func (gm *GraphManager) AddNode(n *Node) *Node {

	//Database

	index, _ := gm.client.Get("nodeCount")
	gm.client.Incr("nodeCount")
	gm.client.Hset("node:"+string(index), "index", []byte(index))

	//Add node to index
	gm.client.Sadd("nodes", index)
	//Local
	n.id = string(index)
	gm.nodes[string(index)] = n
	return n
}

func (gm *GraphManager) DeleteNode(n *Node) {
	//Database
	gm.client.Srem("nodes", []byte(n.id))
	//Local
	gm.nodes[n.id] = nil
	n = nil
}

func (gm *GraphManager) FindNode(name string) *Node {
	return nil
}

func (gm *GraphManager) UpdateNode(n *Node) bool {
	return true
}

func (gm *GraphManager) GetNode(index string) *Node {
	nodeIdx, err := gm.client.Hget("node:"+index, "index")

	if err != nil {
		return nil
	}

	node := new(Node)
	node.id = string(nodeIdx)
	gm.nodes[node.id] = node
	return node
}

//func (gm *GraphManager) GetAdjPairs(node *Node) *[]AdjPair {}

//Add neigbhbor both locally and in db
func (gm *GraphManager) AddNeighbor(node *Node, edge *Edge) {

}

func (gm *GraphManager) GetNeighbors(node *Node) []*Node{

}

//EDGE MANAGEMENT

func (gm *GraphManager) CreateEdge() *Edge {
	return nil
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

func (gm *GraphManager) ClearAll() {
	gm.client.Flushdb()
	gm.nodes = nil
	gm.edges = nil
	gm.client = nil
}

//PRIVATE FUNCTIONS

//func ( gm *GraphManager ) addAdjEntry( )
//func ( gm *GraphManager ) removeAdjEntry( )
