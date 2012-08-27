package socialGraph

import "github.com/ericfode/SpiderDB"

//TODO: rename struct vars to lowercase
//TODO: add commit function
type SocialNode struct {
	Id         string
	Pic        string
	ProperName string
	UserName   string
	Email      string
	Bio        string
	Skills     string
	Github     string
	Edges      map[string][]spiderDB.Edge
	GM         spiderDB.GraphBackend
}

func NewSocialNode(pic string, proper string, user string,
	email string, bio string, skills string, github string,
	gm spiderDB.GraphBackend) *SocialNode {
	sn := new(SocialNode)
	sn.GM = gm
	sn.Pic = pic
	sn.ProperName = proper
	sn.UserName = user
	sn.Email = email
	sn.Bio = bio
	sn.Skills = skills
	sn.Github = github
	sn.Edges = make(map[string][]spiderDB.Edge)
	return sn
}

func SocialNodeConst(id string, GM spiderDB.GraphBackend) spiderDB.Node {
	sn := new(SocialNode)
	sn.SetGM(GM)
	sn.SetID(id)
	return sn
}

func (n *SocialNode) SetGM(gm spiderDB.GraphBackend) {
	n.GM = gm
}

// has this spiderDB.Node been added to the gm?
func (n *SocialNode) IsReged() bool {
	return true
}

func (n *SocialNode) GetID() string {
	return n.Id
}

func (n *SocialNode) SetID(id string) {
	n.Id = id
	/*if n.IsReged() {
		n.GM.UpdateNodeProp(n, "Id", []byte(n.Id))
	}*/
}

func (n *SocialNode) GetPic() string {
	return n.Pic
}

func (n *SocialNode) SetPic(pic string) {
	n.Pic = pic
}

func (n *SocialNode) GetProperName() string {
	return n.ProperName
}

func (n *SocialNode) SetProperName(proper string) {
	n.ProperName = proper
}

func (n *SocialNode) GetUserName() string {
	return n.UserName
}

func (n *SocialNode) SetUserName(name string) {
	n.UserName = name
	/*if n.IsReged() {
		n.GM.UpdateNodeProp(n, "Name", []byte(name))
	}*/
}

func (n *SocialNode) GetEmail() string {
	return n.Email
}

func (n *SocialNode) SetEmail(email string) {
	n.Email = email
	/*if n.IsReged() {
		n.GM.UpdateNodeProp(n, "Email", []byte(email))
	}*/
}

func (n *SocialNode) GetBio() string {
	return n.Bio
}

func (n *SocialNode) SetBio(bio string) {
	n.Bio = bio
}

//User Function
//use bytes package instead for better performace...
/*func (n *SocialNode) SetAwesomeness(awe int) {
	n.Awesomeness = awe
	if n.IsReged() {
		n.GM.UpdateNodeProp(n, "Name", []byte(strconv.Itoa(awe)))
	}
}
*/
//DB only function

func (n *SocialNode) GetSkills() string {
	return n.Skills
}

func (n *SocialNode) SetSkills(skillz string) {
	n.Skills = skillz
}

func (n *SocialNode) GetGit() string {
	return n.Github
}

func (n *SocialNode) SetGit(github string) {
	n.Github = github
}

func (n *SocialNode) SetEdges(edges []spiderDB.Edge) {
	n.Edges = make(map[string][]spiderDB.Edge)
	n.AddEdges(edges)
}

func (n *SocialNode) AddEdge(edge spiderDB.Edge) {
	if n.Edges == nil {
		n.Edges = make(map[string][]spiderDB.Edge)
	}
	n.Edges[edge.GetType()] = append(n.Edges[edge.GetType()], edge)
}

//DB only function
func (n *SocialNode) AddEdges(edges []spiderDB.Edge) {
	for _, edge := range edges {
		n.Edges[edge.GetType()] = append(n.Edges[edge.GetType()], edge)
	}
}

func (n *SocialNode) RemoveEdges(edges []spiderDB.Edge) {
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
		"Id":         []byte(n.Id),
		"Pic":        []byte(n.Pic),
		"ProperName": []byte(n.ProperName),
		"UserName":   []byte(n.UserName),
		"Email":      []byte(n.Email),
		"Bio":        []byte(n.Bio),
		"Skills":     []byte(n.Skills),
		"Github":     []byte(n.Github)}
	return propMap
}

func (n *SocialNode) SetPropMap(props map[string][]byte) {
	n.Id = string(props["Id"])
	n.Pic = string(props["Pic"])
	n.ProperName = string(props["ProperName"])
	n.UserName = string(props["UserName"])
	n.Email = string(props["Email"])
	n.Bio = string(props["Bio"])
	n.Skills = string(props["Skills"])
	n.Github = string(props["Github"])
}

func (n *SocialNode) Equals(other spiderDB.Node) bool {
	if a, ok := other.(*SocialNode); ok {
		if n.GetID() == a.GetID() &&
			n.GetPic() == a.GetPic() &&
			n.GetProperName() == a.GetProperName() &&
			n.GetUserName() == a.GetUserName() &&
			n.GetEmail() == a.GetEmail() &&
			n.GetBio() == a.GetBio() &&
			n.GetSkills() == a.GetSkills() &&
			n.GetGit() == a.GetGit() {
			return true
		}
	}
	return false
}
