package spiderDB

type Edge interface {
	GetID() string
	SetID(id string)

	GetPropMap() map[string][]byte
	SetPropMap(props map[string][]byte)

	IsDirected() bool

	GetType() string

	GetFirstNode() Node
	GetSecondNode() Node

	SetFirstNode(node Node)
	SetSecondNode(node Node)
	Equals(o Edge) bool
}

type EdgeConstructor func(id string, gm GraphBackend) Edge
