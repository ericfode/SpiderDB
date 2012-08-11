package spiderDB

import "fmt"

type EdgeNotAddedToDBError Edge

type NodeNotAddedToDBError Node

type KeyNotFoundError string

type dbError string

func (e dbError) Error() string {
	return fmt.Sprintf("%s", e)
}
