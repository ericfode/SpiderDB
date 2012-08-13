package spiderDB

import "testing"

var gm *GraphManager

func initTestEdges(gm GraphBackend) []*SocialEdge {
	var testEdges = []*SocialEdge{
		&SocialEdge{weight: 43, typ: "knows", GM: gm},
		&SocialEdge{weight: 110, typ: "likes", GM: gm},
		&SocialEdge{weight: 79, typ: "hates", GM: gm},
		&SocialEdge{weight: 2, typ: "stalks", GM: gm},
		&SocialEdge{weight: 53, typ: "knows", GM: gm},
		&SocialEdge{weight: 89, typ: "likes", GM: gm},
		&SocialEdge{weight: 12, typ: "hates", GM: gm},
		&SocialEdge{weight: 99, typ: "stalks", GM: gm},
	}
	return testEdges
}

func initTestNodes(gm GraphBackend) []*SocialNode {
	var testNodes = []*SocialNode{
		&SocialNode{Name: "Bill", Email: "bill@billisAwsome.com", Awesomeness: 40, GM: gm},
		&SocialNode{Name: "Jane", Email: "jane@think.com", Awesomeness: 40, GM: gm},
		&SocialNode{Name: "Sue", Email: "Sue@isueyou.com", Awesomeness: 3240, GM: gm},
		&SocialNode{Name: "Sally", Email: "smadfs@gmail.com", Awesomeness: 30, GM: gm},
		&SocialNode{Name: "Tom", Email: "rawr@hackerschool.com", Awesomeness: 5120, GM: gm},
		&SocialNode{Name: "Domnick", Email: "affiliate@iscamyou.com", Awesomeness: 52, GM: gm},
		&SocialNode{Name: "Eric", Email: "eric@gmail.com", Awesomeness: 52340, GM: gm},
		&SocialNode{Name: "Sarah", Email: "sarah@yahoo.com", Awesomeness: 5546, GM: gm},
		&SocialNode{Name: "Nathan", Email: "shortemail@ineedemail.com", Awesomeness: 43, GM: gm},
		&SocialNode{Name: "That Guy", Email: "anothertroll@myemailwastaken.com", Awesomeness: 51, GM: gm},
		&SocialNode{Name: "That Girl", Email: "troll@girls.com", Awesomeness: 51, GM: gm},
		&SocialNode{Name: "Ugg", Email: "mrr@complain.com", Awesomeness: 5234, GM: gm},
	}
	return testNodes
}
func TestAddSingleNode(t *testing.T) {
	gm = new(GraphManager)
	gm.Initialize()
	defer gm.ClearAll()
	nodes := initTestNodes(gm)

	gm.AddNode(nodes[0])

	t.Logf("node index : %s", nodes[0].GetID())

	if nodes[0].GetID() == "" {
		t.Errorf("id was nil on AddNode")
		return
	}
}

func TestAddNodes(t *testing.T) {

	gm = new(GraphManager)
	gm.Initialize()
	defer gm.ClearAll()
	n := initTestNodes(gm)

	gm.AddNode(n[0])
	gm.AddNode(n[1])
	gm.AddNode(n[2])
	gm.AddNode(n[3])

	t.Logf("node indices : %s %s %s %s", n[0].GetID(), n[1].GetID(), n[2].GetID(), n[3].GetID())

	if n[0].GetID() != "0" {
		t.Errorf("id 0 was %s", n[0].GetID())
		return
	}
	if n[1].GetID() != "1" {
		t.Errorf("id 1 was %s", n[1].GetID())
		return
	}
	if n[2].GetID() != "2" {
		t.Errorf("id 2 was %s", n[2].GetID())
		return
	}
	if n[3].GetID() != "3" {
		t.Errorf("id 3 was %s", n[3].GetID())
		return
	}

}

func TestClear(t *testing.T) {
	gm = new(GraphManager)
	gm.Initialize()
	gm.ClearAll()
	if gm.nodes != nil || gm.edges != nil || gm.client != nil {
		t.Error("GraphManager did not ClearAll")
	}
}

func TestDeleteNode(t *testing.T) {
	gm = new(GraphManager)
	gm.Initialize()
	defer gm.ClearAll()
	nodes := initTestNodes(gm)
	n := nodes[0]
	gm.AddNode(n)
	index := n.GetID()

	gm.DeleteNode(n)

	if nDb, err := gm.GetNode(index, SocialNodeConst); (err != nil) && (nDb == nil) {
		t.Log(err.Error())
	} else if nDb != nil {
		t.Errorf("found node id: %+v", nDb)
		t.Error("GraphManager did not delete node")
		return
	}

}

func TestNodeConstructor(t *testing.T) {
	gm = new(GraphManager)
	gm.Initialize()
	defer gm.ClearAll()
	node := SocialNodeConst("42", gm)
	if node == nil {
		t.Error("Node is nil")
	}
}

func TestGetNode(t *testing.T) {
	gm = new(GraphManager)
	gm.Initialize()
	defer gm.ClearAll()
	nodes := initTestNodes(gm)
	n := nodes[0]
	gm.AddNode(n)
	index := n.GetID()
	nDb, err := gm.GetNode(index, SocialNodeConst)
	if err != nil {
		t.Error(err.Error())
		return
	}
	if nDb == nil {
		t.Errorf("WTF nDb is nil in TestGetNode")
		return
	}
	if nDb.GetID() == "0" {
		t.Error("GraphManager did not get node correctly")
		return
	}

}

func TestAddSingleEdge(t *testing.T) {
	gm = new(GraphManager)
	gm.Initialize()
	defer gm.ClearAll()

	e := initTestEdges(gm)
	edge := e[0]

	gm.AddEdge(edge)

	if edge.GetID() != "0" {
		t.Errorf("GraphManager: id 0 was %d", edge.GetID())
	}
}

func TestAddEdges(t *testing.T) {
	gm = new(GraphManager)
	gm.Initialize()
	defer gm.ClearAll()

	e := initTestEdges(gm)

	gm.AddEdge(e[0])
	gm.AddEdge(e[1])
	gm.AddEdge(e[2])
	gm.AddEdge(e[3])

	if e[0].GetID() != "0" {
		t.Errorf("edge id 0 was %d", e[0].GetID())
		return
	}
	if e[1].GetID() != "1" {
		t.Errorf("edge id 1 was %d", e[1].GetID())
		return
	}
	if e[2].GetID() != "2" {
		t.Errorf("edge id 2 was %d", e[2].GetID())
		return
	}
	if e[3].GetID() != "3" {
		t.Errorf("edge id 3 was %d", e[3].GetID())
		return
	}

}
