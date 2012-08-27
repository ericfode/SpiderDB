package socialGraph

import "github.com/ericfode/SpiderDB"

type MessageNode struct {
	id    string
	text  string
	Edges map[string][]spiderDB.Edge
	GM    spiderDB.GraphBackend
}

func NewMessageNode(text string) *MessageNode {
	mn := new(MessageNode)
	mn.text = text
	mn.Edges = make(map[string][]spiderDB.Edge)
	return mn
}
func MessageNodeConst(id string, GM spiderDB.GraphBackend) spiderDB.Node {
	mn := new(MessageNode)
	mn.Edges = make(map[string][]spiderDB.Edge)
	mn.SetGM(GM)
	mn.SetID(id)
	return mn
}
func (n *MessageNode) GetID() string {
	return n.id
}
func (n *MessageNode) SetID(id string) {
	n.id = id
}
func (n *MessageNode) GetText() string {
	return n.text
}
func (n *MessageNode) SetText(msg string) {
	n.text = msg
}
func (n *MessageNode) SetEdges(edges []spiderDB.Edge) {
	n.Edges = make(map[string][]spiderDB.Edge)
	n.AddEdges(edges)
}
func (n *MessageNode) AddEdge(edge spiderDB.Edge) {
	if n.Edges == nil {
		n.Edges = make(map[string][]spiderDB.Edge)
	}
	n.Edges[edge.GetType()] = append(n.Edges[edge.GetType()], edge)
}
func (n *MessageNode) AddEdges(edges []spiderDB.Edge) {
	for _, edge := range edges {
		n.Edges[edge.GetType()] = append(n.Edges[edge.GetType()], edge)
	}
}
func (n *MessageNode) RemoveEdges(edges []spiderDB.Edge) {
	for _, deleteMe := range edges {
		for id, checkMe := range n.Edges[deleteMe.GetType()] {
			if checkMe.GetID() == deleteMe.GetID() {
				n.Edges[deleteMe.GetType()][id] = nil
			}
		}
	}
}
func (n *MessageNode) GetPropMap() map[string][]byte {
	var propMap = map[string][]byte{
		"text": []byte(n.GetText())}
	return propMap
}
func (n *MessageNode) SetPropMap(props map[string][]byte) {
	n.text = string(props["text"])
}
func (n *MessageNode) SetGM(gm spiderDB.GraphBackend) {
	n.GM = gm
}
func (n *MessageNode) Equals(other spiderDB.Node) bool {
	oth, ok := other.(*MessageNode)

	if !ok {
		return false
	}

	if n.GetID() == oth.GetID() &&
		n.GetText() == oth.GetText() {
		return true
	}
	return false
}
