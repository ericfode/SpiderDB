package spiderDB

type GraphBackend interface {
	Initialize()
	Connect(port int, id string, db int)
	AddNode(n Node)
	DeleteNode(n Node)
	FindNode(string, NodeConstructor) (Node, error)
	UpdateNode(n Node) error
	UpdateNodeProp(n Node, prop string, value []byte) error
	GetNode(string, NodeConstructor) (Node, error)
	Attach(node1 Node, node2 Node, edge Edge)
	GetNeighbors(node Node) ([]Connection)
	AddEdge(e Edge)
	DeleteEdge(e Edge)
	FindEdge(id int) Edge
	UpdateEdge(e Edge) bool
	UpdateEdgeProp(e Edge, prop string, value []byte) error
	GetEdge(id int) Edge
	GetNodeEdges(n Node) map[string][]Edge
	ClearAll()
}
