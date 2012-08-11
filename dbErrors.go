package spiderDB

import "fmt"

type EdgeNotAddedToDBError *Edge

type NodeNotAddedToDBError *Node

type KeyNotFoundError string

type dbError string

func (e dbError) Error() string {
	return fmt.Sprintf("%s", e)
}

func (e KeyNotFoundError) Error() string {
	return fmt.Sprintf("%s", e)
}

func (e *EdgeNotAddedToDBError) Error() string {
	return fmt.Sprintf("Edge %v has not been added to the db, run AddEdge first", e)
}

func (e *NodeNotAddedToDBError) Error() string {
	return fmt.Sprintf("Node %v has not been added to the db, run AddNode first", e)
}
