package spiderDB

import "strconv"

type SocialEdge struct {
	id      int
	weight  int
	typ     string
	GM      GraphBackend
	fstNode Node
	sndNode Node
}

func SocialEdgeConst(id string) Edge {
	edge := new(SocialEdge)
	edge.SetID(id)
	return edge
}

func (s *SocialEdge) GetID() string {
	return strconv.Itoa(s.id)
}

//TODO: prolly should be having db do this....
func (s *SocialEdge) SetID(id string) {
	s.id, _ = strconv.Atoi(id)
}

func (s *SocialEdge) GetWeight() int {
	return s.weight
}

func (s *SocialEdge) SetWeight(weight string) {
	s.weight, _ = strconv.Atoi(weight)
}
func (s *SocialEdge) GetType() string {
	return s.typ
}
func (s *SocialEdge) SetType(typestr string) {

}

func (s *SocialEdge) SetPropMap(props map[string][]byte) {
	s.id = BytesToInt(props["Id"])
	s.weight = BytesToInt(props["Weight"])
	s.typ = string(props["Type"])
}

func (s *SocialEdge) GetPropMap() map[string][]byte {
	var propMap = map[string][]byte{
		"Id":     IntToBytes(s.id),
		"Weight": IntToBytes(s.weight),
		"Type":   []byte(s.typ)}
	return propMap
}

func (s *SocialEdge) GetFirstNode() Node {
	return s.fstNode
}
func (s *SocialEdge) GetSecondNode() Node {
	return s.sndNode
}
func (s *SocialEdge) SetFirstNode(node Node) {
	s.fstNode = node
}
func (s *SocialEdge) SetSecondNode(node Node) {
	s.sndNode = node
}

func (s *SocialEdge) IsDirected() bool {
	return false
}
