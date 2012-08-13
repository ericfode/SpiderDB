package spiderDB

import "strconv"

//TODO: rename struct vars to lowercase
type SocialNode struct {
	Id          string
	Name        string
	Email       string
	Awesomeness int
	Edges       map[string][]Edge
	GM          GraphBackend
}

func NewSocialNode(name string, email string, awe int, gm GraphBackend) *SocialNode {
	sn := new(SocialNode)
	sn.Awesomeness = awe
	sn.Name = name
	sn.Email = email
	sn.GM = gm
	return sn
}

func SocialNodeConst(id string) Node{
	sn := new(SocialNode)
	sn.SetID(id)
	return sn
}

func (n *SocialNode) SetGM(gm GraphBackend) {
	n.GM = gm
}

// has this node been added to the gm?
func (n *SocialNode) IsReged() bool {
	return true
}

func (n *SocialNode) GetID() string {
	return n.Id
}

func (n *SocialNode) SetID(id string) {
	n.Id = id
	if n.IsReged() {
		n.GM.UpdateNodeProp(n, "Id", []byte(n.Id))
	}
}

func (n *SocialNode) GetName() string {
	return n.Name
}

func (n *SocialNode) SetName(name string) {
	n.Name = name
	if n.IsReged() {
		n.GM.UpdateNodeProp(n, "Name", []byte(name))
	}
}

func (n *SocialNode) GetEmail() string {
	return n.Email
}

func (n *SocialNode) SetEmail(email string) {
	n.Email = email
	if n.IsReged() {
		n.GM.UpdateNodeProp(n, "Email", []byte(email))
	}
}

func (n *SocialNode) GetAwesomeness() int {
	return n.Awesomeness
}

//User Function
//use bytes package instead for better performace...
func (n *SocialNode) SetAwesomeness(awe int) {
	n.Awesomeness = awe
	if n.IsReged() {
		n.GM.UpdateNodeProp(n, "Name", []byte(strconv.Itoa(awe)))
	}
}

//DB only function
func (n *SocialNode) SetEdges(edges map[string][]Edge) {
	n.Edges = edges
}

//DB only function
func (n *SocialNode) AddEdges(edges []Edge) {
	//	for _, edge := range edges {
	//		append(n.Edges[edge.GetType()], edge)
	//	}
}

func (n *SocialNode) RemoveEdge(edges []Edge) {
	for _, findedge := range edges {
		for index, edge := range n.Edges[findedge.GetType()] {
			if edge.GetID() == findedge.GetID() {
				n.Edges[findedge.GetType()][index] = nil
			}
		}
	}
}

func (n *SocialNode) GetPropMap() map[string][]byte {
	var propMap = map[string][]byte{
		"Id":          []byte(n.Id),
		"Name":        []byte(n.Name),
		"Email":       []byte(n.Email),
		"Awesomeness": []byte(strconv.Itoa(n.Awesomeness))}
	return propMap
}

func (n *SocialNode) SetPropMap(props map[string][]byte) {
	n.Id = string(props["Id"])
	n.Name = string(props["Name"])
	n.Email = string(props["Email"])
	n.Awesomeness, _ = strconv.Atoi(string(props["Awesomeness"]))
}
