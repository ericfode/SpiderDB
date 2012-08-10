package main

//TO-DO:
//  + make type keys and type set
//  + make sorted set of all nodes with type id as score
//  + when inserting node, add to sorted set 
//  + add register/deregister type functions
//      + deregister should only work if no nodes of that type exist
//      + register should only work if type does not already exist

import (
	"github.com/hoisie/redis.go"
)

//Initializes database at startup

var db redis.Client

type GraphManager struct {
	nodes map[int]*Node
	edges map[int]*Edge
}

func (gm *GraphManager) Initialize() {
	db.Set("nodeCount", 0)
}

func (gm *GraphManager) Connect(port int, ipAddr String, dbNum int) {
	db.Addr = ipAddr + ":" + port
	db.Db = dbNum
}

// NODE MANAGEMENT

func (gm *GraphManager) AddNode(n *Node) *Node {

	//Database

	index := db.Get("nodeCount")
	db.Incr("nodeCount")
	db.Hset("node:"+index, "index", []byte(index))
	
	//Add node to index
	db.Sadd("nodes",[]byte(index))
	//Local
	n.id = index
	gm.nodes[index] = n
	return n
}

func (gm *GraphManager) DeleteNode(n *Node) {
	//Database
	db.SRem("nodes",n.id)
	//Local
	gm.nodes[n.id] = nil
	n = nil
}


func (gm *GraphManager) FindNode(name String) *Node {

}

func (gm *GraphManger) UpdateNode(n *Node) Bool {

}

func (gm *GraphManger) GetNode(Id String) *Node {
	db.Get("node:"+index)
}

func (gm *GraphManager) GetAdjPairs(node *Node) *[]AdjPair {

}

//Add neigbhbor both locally and in db
func (gm *GraphManager) AddNeighbor(node *Node, edge *Edge) {

}

//EDGE MANAGEMENT

func (gm *GraphManager) CreateEdge() *Edge {

}

func (gm *GraphManager) DeleteEdge(e *Edge) {

}

func (gm *GraphManager) FindEdge(id int) *Edge {

}

func (gm *GraphManger) UpdateEdge(e *Edge) Bool {

}

func (gm *GraphManger) GetEdge(id int) *Edge {

}

//PRIVATE FUNCTIONS

//func ( gm *GraphManager ) addAdjEntry( )
//func ( gm *GraphManager ) removeAdjEntry( )
