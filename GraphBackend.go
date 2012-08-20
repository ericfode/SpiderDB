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
	GetNeighbors(node Node, constE EdgeConstructor, constN NodeConstructor) ([]Connection, error)
	AddEdge(e Edge)
	DeleteEdge(e Edge)
	FindEdge(id string, construct EdgeConstructor) (Edge, error)
	UpdateEdge(e Edge) error
	UpdateEdgeProp(e Edge, prop string, value []byte) error
	GetEdge(id string, construct EdgeConstructor) (Edge, error)
	GetOutgoingNodeEdges(n Node, construct EdgeConstructor) []Edge
	GetIncomingNodeEdges(n Node, construct EdgeConstructor) []Edge
	GetAllNodeEdges(n Node, construct EdgeConstructor) []Edge
	NodeFromHash([][]byte, NodeConstructor) (Node, bool)
	EdgeFromHash([][]byte, EdgeConstructor) (Edge, bool)

	ClearAll()
}
