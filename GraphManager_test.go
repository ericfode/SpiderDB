package spiderDB

import "testing"

var gm *GraphManager

func TestAddNode(t *testing.T) {

	gm = new(GraphManager)
	gm.Initialize()

	n := new(Node)
	gm.AddNode(n)

	t.Logf("node index : %s", n.id)

	if n.id == "" {
		t.Errorf("id was nil on AddNode")
	}
	gm.ClearAll()
}

func TestAddNodes(t *testing.T) {

	gm = new(GraphManager)
	gm.Initialize()

	n := new(Node)
	n1 := new(Node)
	n2 := new(Node)
	n3 := new(Node)

	gm.AddNode(n)
	gm.AddNode(n1)
	gm.AddNode(n2)
	gm.AddNode(n3)

	t.Logf("node indices : %s %s %s %s", n.id, n1.id, n2.id, n3.id)

	if n.id == "" {
		t.Errorf("id 0 was %s", n.id)
	}
	if n1.id != "1" {
		t.Errorf("id 1 was %s", n1.id)
	}
	if n2.id != "2" {
		t.Errorf("id 2 was %s", n2.id)
	}
	if n3.id != "3" {
		t.Errorf("id 3 was %s", n3.id)
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
