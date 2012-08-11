package spiderDB

import "testing"

var gm *GraphManager

func initTestNodes() []*SocialNode {
	var testNodes = []*SocialNode{
		&SocialNode{Name: "Bill", Email: "bill@billisAwsome.com", Awesomeness: 40},
		&SocialNode{Name: "Jane", Email: "jane@think.com", Awesomeness: 40},
		&SocialNode{Name: "Sue", Email: "Sue@isueyou.com", Awesomeness: 3240},
		&SocialNode{Name: "Sally", Email: "smadfs@gmail.com", Awesomeness: 30},
		&SocialNode{Name: "Tom", Email: "rawr@hackerschool.com", Awesomeness: 5120},
		&SocialNode{Name: "Domnick", Email: "affiliate@iscamyou.com", Awesomeness: 52},
		&SocialNode{Name: "Eric", Email: "eric@gmail.com", Awesomeness: 52340},
		&SocialNode{Name: "Sarah", Email: "sarah@yahoo.com", Awesomeness: 5546},
		&SocialNode{Name: "Nathan", Email: "shortemail@ineedemail.com", Awesomeness: 43},
		&SocialNode{Name: "That Guy", Email: "anothertroll@myemailwastaken.com", Awesomeness: 51},
		&SocialNode{Name: "That Girl", Email: "troll@girls.com", Awesomeness: 51},
		&SocialNode{Name: "Ugg", Email: "mrr@complain.com", Awesomeness: 5234},
	}
	return testNodes
}
func TestAddSingleNode(t *testing.T) {
	nodes := initTestNodes()
	gm = new(GraphManager)
	gm.Initialize()

	gm.AddNode(nodes[0])

	t.Logf("node index : %s", nodes[0].GetID())

	if nodes[0].GetID() == "" {
		t.Errorf("id was nil on AddNode")
	}
	gm.ClearAll()
}

func TestAddNodes(t *testing.T) {

	gm = new(GraphManager)
	gm.Initialize()
	n := initTestNodes()

	gm.AddNode(n[0])
	gm.AddNode(n[1])
	gm.AddNode(n[2])
	gm.AddNode(n[3])

	t.Logf("node indices : %s %s %s %s", n[0].GetID(), n[1].GetID(), n[2].GetID(), n[3].GetID())

	if n[0].GetID() == "" {
		t.Errorf("id 0 was %s", n[0].GetID())
	}
	if n[1].GetID() != "1" {
		t.Errorf("id 1 was %s", n[1].GetID())
	}
	if n[2].GetID() != "2" {
		t.Errorf("id 2 was %s", n[2].GetID())
	}
	if n[3].GetID() != "3" {
		t.Errorf("id 3 was %s", n[3].GetID())
	}

	gm.ClearAll()
}

func TestClear(t *testing.T) {
	gm.Initialize()
	gm.ClearAll()
	if gm.nodes != nil || gm.edges != nil || gm.client != nil {
		t.Error("GraphManager did not ClearAll")
	}
}

func TestDeleteNode(t *testing.T) {
	gm.Initialize()
	nodes := initTestNodes()
	n := nodes[0]
	gm.AddNode(n)
	index := n.GetID()

	gm.DeleteNode(n)

	nDb := gm.GetNode(index)

	if nDb != nil {
		t.Errorf("found node id: %+v", nDb)
		t.Error("GraphManager did not delete node")
	}

	gm.ClearAll()
}

func TestGetNode(t *testing.T) {
	gm.Initialize()
	nodes := initTestNodes()
	n := nodes[0]
	gm.AddNode(n)
	index := n.GetID()

	nDb := gm.GetNode(index)

	if nDb.GetID() != n.GetID() {
		t.Error("GraphManager did not get node correctly")
	}

	gm.ClearAll()
}

/*
func TestAddEdge(t *testing.T) {
	gm.Initialize()
	e := new(Edge)
	gm.AddEdge(e)

	if e.GetID() != "0" {
		t.Errorf("GraphManager: id 0 was %s", e.GetID())
	}

	gm.ClearAll()
}
*/
