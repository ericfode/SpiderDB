package spiderDB

import "strconv"

type SocialEdge struct {
	id     int
	weight int
	typ    string
	GM     *GraphBackend
}

func (s *SocialEdge) GetId() string {
	return strconv.Itoa(s.id)
}
func (s *SocialEdge) SetId(id string) {

}
func (s *SocialEdge) GetWeight() int {
	return s.weight
}
func (s *SocialEdge) SetWeight(weight int) {

}
func (s *SocialEdge) GetType() string {
	return s.typ
}
func (s *SocialEdge) SetType(typestr string) {

}
