package socialGraph

import "strconv"
import "github.com/ericfode/SpiderDB"

type SocialEdge struct {
	id      int
	typ     string
	date    int
	GM      spiderDB.GraphBackend
	fstNode spiderDB.Node
	sndNode spiderDB.Node
}

func NewSocialEdge(date int, typ string, gm spiderDB.GraphBackend) *SocialEdge {
	edge := new(SocialEdge)
	edge.SetGM(gm)
	edge.SetType(typ)
	edge.SetDate(date)
	return edge
}

func SocialEdgeConst(id string, GM spiderDB.GraphBackend) spiderDB.Edge {
	edge := new(SocialEdge)
	edge.SetGM(GM)
	edge.SetID(id)
	return edge
}

func (e *SocialEdge) SetGM(gm spiderDB.GraphBackend) {
	e.GM = gm
}

func (s *SocialEdge) GetID() string {
	return strconv.Itoa(s.id)
}

//TODO: prolly should be having db do this....
func (s *SocialEdge) SetID(id string) {
	s.id, _ = strconv.Atoi(id)
}

func (s *SocialEdge) GetDate() int {
	return s.date
}

func (s *SocialEdge) SetDate(date int) {
	s.date = date
}
func (s *SocialEdge) GetType() string {
	return s.typ
}
func (s *SocialEdge) SetType(typestr string) {
	s.typ = typestr
}

func (s *SocialEdge) SetPropMap(props map[string][]byte) {
	s.id = spiderDB.BytesToInt(props["Id"])
	s.date = spiderDB.BytesToInt(props["Date"])
	s.typ = string(props["Type"])
}

func (s *SocialEdge) GetPropMap() map[string][]byte {
	var propMap = map[string][]byte{
		"Id":   spiderDB.IntToBytes(s.id),
		"Date": spiderDB.IntToBytes(s.date),
		"Type": []byte(s.typ)}
	return propMap
}

func (s *SocialEdge) GetFirstNode() spiderDB.Node {
	return s.fstNode
}
func (s *SocialEdge) GetSecondNode() spiderDB.Node {
	return s.sndNode
}
func (s *SocialEdge) GetOtherNode(node spiderDB.Node) spiderDB.Node {
	if s.fstNode.GetID() == node.GetID() {
		return s.sndNode
	}
	if s.sndNode.GetID() == node.GetID() {
		return s.fstNode
	}
	return nil
	//TODO : Make appropriately descriptive error
}
func (s *SocialEdge) SetFirstNode(node spiderDB.Node) {
	s.fstNode = node
}
func (s *SocialEdge) SetSecondNode(node spiderDB.Node) {
	s.sndNode = node
}

func (s *SocialEdge) IsDirected() bool {
	return true
}

func (s *SocialEdge) Equals(o spiderDB.Edge) bool {
	if a, ok := o.(*SocialEdge); ok {
		if (s.GetID() == a.GetID() &&
			s.GetDate() == a.GetDate() &&
			s.GetType() == a.GetType()){
			return true
		}
	}
	return false
}
