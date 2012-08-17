package spiderDB_test

import "testing"
import "github.com/HackerSchool12/SpiderDB"
import "github.com/HackerSchool12/SpiderDB/socialGraph"

//TODO: update tests to use SpiderDB.Node.Equals instead of just compairing id
var gm *spiderDB.GraphManager

func initTestEdges(gm spiderDB.GraphBackend) []*socialGraph.SocialEdge {
	var testEdges = []*socialGraph.SocialEdge{
		socialGraph.NewSocialEdge(43, "knows", gm),
		socialGraph.NewSocialEdge(110, "likes", gm),
		socialGraph.NewSocialEdge(79, "hates", gm),
		socialGraph.NewSocialEdge(2, "stalks", gm),
		socialGraph.NewSocialEdge(53, "knows", gm),
		socialGraph.NewSocialEdge(89, "likes", gm),
		socialGraph.NewSocialEdge(12, "hates", gm),
		socialGraph.NewSocialEdge(99, "stalks", gm),
	}
	return testEdges
}

func initTestNodes(gm spiderDB.GraphBackend) []*socialGraph.SocialNode {
	var testNodes = []*socialGraph.SocialNode{
		socialGraph.NewSocialNode("Joe", "joe@joe.com", 120, gm),
		socialGraph.NewSocialNode("Bill", "bill@billisAwsome.com", 40, gm),
		socialGraph.NewSocialNode("Jane", "jane@think.com", 40, gm),
		socialGraph.NewSocialNode("Sue", "Sue@isueyou.com", 3240, gm),
		socialGraph.NewSocialNode("Sally", "smadfs@gmail.com", 30, gm),
		socialGraph.NewSocialNode("Tom", "rawr@hackerschool.com", 5120, gm),
		socialGraph.NewSocialNode("Domnick", "affiliate@iscamyou.com", 52, gm),
		socialGraph.NewSocialNode("Eric", "eric@gmail.com", 52340, gm),
		socialGraph.NewSocialNode("Sarah", "sarah@yahoo.com", 5546, gm),
		socialGraph.NewSocialNode("Nathan", "shortemail@ineedemail.com", 43, gm),
		socialGraph.NewSocialNode("That Guy", "anothertroll@myemailwastaken.com", 51, gm),
		socialGraph.NewSocialNode("That Girl", "troll@girls.com", 51, gm),
		socialGraph.NewSocialNode("Ugg", "mrr@complain.com", 5234, gm),
	}
	return testNodes
}
func TestAddSingleNode(t *testing.T) {
	gm = new(spiderDB.GraphManager)
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

	gm = new(spiderDB.GraphManager)
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
	gm = new(spiderDB.GraphManager)
	gm.Initialize()
	gm.ClearAll()
	//if gm.nodes != nil || gm.edges != nil || gm.client != nil {
	//	t.Error("GraphManager did not ClearAll")
	//}
}

func TestDeleteNode(t *testing.T) {
	gm = new(spiderDB.GraphManager)
	gm.Initialize()
	defer gm.ClearAll()
	nodes := initTestNodes(gm)
	n := nodes[0]
	gm.AddNode(n)
	index := n.GetID()

	gm.DeleteNode(n)

	if nDb, err := gm.GetNode(index, socialGraph.SocialNodeConst); (err != nil) && (nDb == nil) {
		if _, ok := err.(*spiderDB.KeyNotFoundError); ok {
			return
		} else {
			t.Error("GM failed to delete node and err properly")
		}
	} else if nDb != nil {
		t.Errorf("found node id: %+v", nDb)
		t.Error("GraphManager did not delete node")
		return
	}

}

func TestNodeConstructor(t *testing.T) {
	gm = new(spiderDB.GraphManager)
	gm.Initialize()
	defer gm.ClearAll()
	node := socialGraph.SocialNodeConst("42", gm)
	if node == nil {
		t.Error("Node is nil")
	}
}

func TestGetNode(t *testing.T) {
	gm = new(spiderDB.GraphManager)
	gm.Initialize()
	defer gm.ClearAll()
	nodes := initTestNodes(gm)
	n := nodes[0]
	gm.AddNode(n)
	index := n.GetID()
	nDb, err := gm.GetNode(index, socialGraph.SocialNodeConst)
	if err != nil {
		t.Error(err.Error())
		return
	}
	if nDb == nil {
		t.Errorf("nDb is nil in TestGetNode")
		return
	}
	if !nDb.Equals(n) {
		t.Errorf(`Saved node and node from GetAllNodes not Equal.../n
					Expected : %v /n
					Actual : %v`, n, nDb)
		return
	}

}

func TestEdgeConstructor(t *testing.T) {
	gm = new(spiderDB.GraphManager)
	gm.Initialize()
	defer gm.ClearAll()

	e := socialGraph.SocialEdgeConst("94", gm)
	if e == nil {
		t.Error("Edge is nil")
	}
}

func TestGetEdge(t *testing.T) {
	gm = new(spiderDB.GraphManager)
	gm.Initialize()
	defer gm.ClearAll()

	edges := initTestEdges(gm)
	e := edges[0]
	gm.AddEdge(e)
	id := e.GetID()

	edge, err := gm.GetEdge(id, socialGraph.SocialEdgeConst)

	if err != nil {
		t.Error(err.Error())
	}
	if edge == nil {
		t.Error("GM did not retrieve edge (nil)")
	}
}

func TestAddSingleEdge(t *testing.T) {
	gm = new(spiderDB.GraphManager)
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
	gm = new(spiderDB.GraphManager)
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

func TestGetAllNodesSingle(t *testing.T) {
	gm = new(spiderDB.GraphManager)
	gm.Initialize()
	defer gm.ClearAll()
	nodes := initTestNodes(gm)
	gm.AddNode(nodes[0])
	allNodes, err := gm.GetAllNodes(socialGraph.SocialNodeConst)
	if err != nil {
		t.Error(err.Error())
	}
	if len(allNodes) != 1 {
		t.Errorf("Unexpected length (should have been 1) was %d", len(allNodes))
		return
	}
	if allNodes[0] == nil {
		t.Error("Got node was nil")
	}
	if !allNodes[0].Equals(nodes[0]) {
		t.Errorf(`Saved node and node from GetAllNodes not Equal.../n
					Expected : %v /n
					Actual : %v`, nodes[0], allNodes[0])
		return
	}
}

func TestGetAllNodesGroup(t *testing.T) {
	gm = new(spiderDB.GraphManager)
	gm.Initialize()
	defer gm.ClearAll()

	nodes := initTestNodes(gm)
	for _, val := range nodes {
		gm.AddNode(val)
	}
	allNodes, err := gm.GetAllNodes(socialGraph.SocialNodeConst)
	if err != nil {
		t.Error(err.Error())
	}
	if len(allNodes) != len(nodes) {
		t.Errorf("len(allNodes) expected: %d\n actual:%d", len(nodes), len(allNodes))
		return
	}
	for i, val := range allNodes {
		if !val.Equals(nodes[spiderDB.StringToInt(val.GetID())]) {
			t.Errorf(`Saved node and node from GetAllNodes not Equals.../n
					Expected : %v /n
					Actual : %v /n
					Index : %d`, nodes[spiderDB.StringToInt(val.GetID())], val, i)
			return
		}
	}
}

func initTestNodeByteAA() [][]byte {
	return [][]byte{[]byte("Name"), []byte("Joe"), []byte("Email"),
		[]byte("joe@joe.com"), []byte("Awesomeness"), []byte("120")}
}

func TestNodeFromHash(t *testing.T) {
	gm = new(spiderDB.GraphManager)
	gm.Initialize()
	defer gm.ClearAll()

	nodes := initTestNodes(gm)
	gm.AddNode(nodes[0])

	nodeFromDB, errDB := gm.FindNode("node:0", socialGraph.SocialNodeConst)
	nodeFromHash, ok := gm.NodeFromHash(initTestNodeByteAA(), socialGraph.SocialNodeConst)

	nodeFromHash.SetID("0")

	if errDB != nil {
		t.Error(errDB)
	}

	if !ok {
		t.Error("Node From Hash failed")
	}

	//compare nodes
	if !nodeFromDB.Equals(nodeFromHash) {
		t.Errorf("Node From Hash incorrect: \ndb: %v \nhs: %v", nodeFromDB, nodeFromHash)
	}
}

func TestGetNodeEdges(t *testing.T) {
	gm = new(spiderDB.GraphManager)
	gm.Initialize()
	defer gm.ClearAll()

	nodes := initTestNodes(gm)
	edges := initTestEdges(gm)

	for _, v := range nodes {
		gm.AddNode(v)
	}
	for _, v := range edges {
		gm.AddEdge(v)
	}

}
