package spiderDB

type Node interface {
	GetID() string //DB
	SetID(id string)

	SetEdges(edges map[string][]Edge)
	AddEdges(edge []Edge)
	RemoveEdge(string []Edge)

	GetPropMap() map[string][]byte
	SetPropMap(props map[string][]byte)

	SetGM(gm GraphBackend)
}

type NodeConstructor func(id string, gm GraphBackend) Node
