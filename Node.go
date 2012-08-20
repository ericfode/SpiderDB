package spiderDB

type Node interface {
	GetID() string //DB
	SetID(id string)

	SetEdges(edges []Edge)
	AddEdge(edge Edge)
	AddEdges(edges []Edge)
	RemoveEdges(edges []Edge)

	GetPropMap() map[string][]byte
	SetPropMap(props map[string][]byte)

	SetGM(gm GraphBackend)

	Equals(other Node) bool
}

type NodeConstructor func(id string, gm GraphBackend) Node
