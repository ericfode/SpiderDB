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
	GetNeighbors(node Node) []Connection
	AddEdge(e Edge)
	DeleteEdge(e Edge)
	FindEdge(id string, construct EdgeConstructor) (Edge, error)
	UpdateEdge(e Edge) error
	UpdateEdgeProp(e Edge, prop string, value []byte) error
	GetEdge(id string, construct EdgeConstructor) (Edge, error)
	GetNodeEdges(n Node) map[string][]Edge
	NodeFromHash([][]byte, NodeConstructor) (Node, bool)
	EdgeFromHash([][]byte, EdgeConstructor) (Edge, bool)

	ClearAll()
}
