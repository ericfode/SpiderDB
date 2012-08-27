package spiderDB_test

import "testing"
import "github.com/Ericfode/SpiderDB"
import "github.com/Ericfode/SpiderDB/socialGraph"

import "fmt"

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
	var testNodes []*socialGraph.SocialNode

	const numdum = 8
	pics := [numdum]string{"picTEST", "http://4.bp.blogspot.com/-Q2hjS1dS1R8/T4YXpOfNjOI/AAAAAAAAAxQ/c-V_1FkMYmo/s1600/Bug.jpg",
		"https://encrypted-tbn1.google.com/images?q=tbn:ANd9GcRs5AS0g3hHRdJsO7gBgwu9v1Hr4grtuc_G1dh59MbxEVW3VH-GNw",
		"https://encrypted-tbn3.google.com/images?q=tbn:ANd9GcSJVzRTk5jiGvRIcKQZs-pm4__kMQOWae0WGGl3H32xZCTvci9U",
		"https://encrypted-tbn3.google.com/images?q=tbn:ANd9GcQ6VCAy3UhBqNohPBG1Dr5nVd2WfwTLnINK_pmh0Wo7RUPh7vwpjw",
		"https://encrypted-tbn1.google.com/images?q=tbn:ANd9GcQ677iObh3n9DhnfwvpFUH-ksX9mXv3kyS_h7npytmLACpe9EZX",
		"https://encrypted-tbn3.google.com/images?q=tbn:ANd9GcR94C_rLFc1arqiV_Dmi6LHIQzEVWvOFJg7TxdpdR-PxtVxVLAr",
		"https://encrypted-tbn3.google.com/images?q=tbn:ANd9GcRPaCON4nIzFMqrfCVuWAn8HoD0zH-ir-KovxFxwgy6ocUlYxHJ"}
	bios := [numdum]string{
		"bioTEST",
		"I don't exist",
		"I am from HERE",
		"I am from THERE",
		"I am from A",
		"I am from B",
		"I am from C",
		"I am from D"}
	names := [numdum]string{"nameTEST", "Wedunno", "Joe", "Bill", "Jane", "Sue", "Sally", "Tom"}
	users := [numdum]string{"userTEST", "Whothatis", "jmk", "bill-o-rama", "sparkles", "user", "user", "uzaaah"}
	github := [numdum]string{"githubTEST",
		"http://github.com/IAMAUSER",
		"http://github.com/IAMAUSER",
		"http://github.com/IAMAUSER",
		"http://github.com/IAMAUSER",
		"http://github.com/IAMAUSER",
		"http://github.com/IAMAUSER",
		"http://github.com/IAMAUSER"}

	for k, _ := range pics {
		newNode := socialGraph.NewSocialNode(pics[k], names[k], users[k], "emailTEST", bios[k], "skillsTEST", github[k], gm)
		testNodes = append(testNodes, newNode)
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
	if !edge.Equals(edges[0]) {
		t.Errorf("Edges are not equal")
	}
}

func TestAddSingleEdge(t *testing.T) {
	gm = new(spiderDB.GraphManager)
	gm.Initialize()
	defer gm.ClearAll()

	e := initTestEdges(gm)
	edge := e[0]

	gm.AddEdge(edge)

	if !edge.Equals(e[0]) {
		t.Errorf("Edges are not equal")
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
	return [][]byte{[]byte("Pic"), []byte("picTEST"),
		[]byte("ProperName"), []byte("nameTEST"),
		[]byte("UserName"), []byte("userTEST"),
		[]byte("Email"), []byte("emailTEST"),
		[]byte("Bio"), []byte("bioTEST"),
		[]byte("Skills"), []byte("skillsTEST"),
		[]byte("Github"), []byte("githubTEST")}
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

func TestGetNeighbors(t *testing.T) {
	gm = new(spiderDB.GraphManager)

	gm.Initialize()
	defer gm.ClearAll()

	hgw := socialGraph.NewSocialNode("http://upload.wikimedia.org/wikipedia/commons/thumb/7/7f/H_G_Wells_pre_1922.jpg/220px-H_G_Wells_pre_1922.jpg",
		"Herbert Wells", "H.G.", "traveler@time.future", "Writer", "SciFi", "whatisgithub?", gm)
	msg := socialGraph.NewMessageNode("Will he ever return?")

	edg := socialGraph.NewSocialEdge(1898, "jittered", gm)

	gm.AddNode(hgw)
	gm.AddNode(msg)
	gm.AddEdge(edg)

	gm.Attach(hgw, msg, edg)

	neighbors, err := gm.GetNeighbors(hgw, socialGraph.SocialEdgeConst, socialGraph.SocialNodeConst)

	if err != nil {
		t.Error(err.Error())
	}

	if len(neighbors) != 1 {
		fmt.Printf("***************%v ************\n", neighbors)
		t.Errorf("GetNeighbors Failed - %v neighbors", len(neighbors))
	}

	if err != nil {
		t.Error(err)
	}
}

func TestMultipleNeighbors(t *testing.T) {
	gm = new(spiderDB.GraphManager)

	gm.Initialize()
	defer gm.ClearAll()

	hgw := socialGraph.NewSocialNode("http://upload.wikimedia.org/wikipedia/commons/thumb/7/7f/H_G_Wells_pre_1922.jpg/220px-H_G_Wells_pre_1922.jpg",
		"Herbert Wells", "H.G.", "traveler@time.future", "Writer", "SciFi", "whatisgithub?", gm)
	msg := socialGraph.NewMessageNode("Will he ever return?")
	msg1 := socialGraph.NewMessageNode("look, it's another message!")
	msg2 := socialGraph.NewMessageNode("yet another message!")
	edg := socialGraph.NewSocialEdge(1898, "jittered", gm)
	edg1 := socialGraph.NewSocialEdge(1899, "jittered", gm)
	edg2 := socialGraph.NewSocialEdge(1900, "jittered", gm)

	gm.AddNode(hgw)
	gm.AddNode(msg)
	gm.AddNode(msg1)
	gm.AddNode(msg2)
	gm.AddEdge(edg)
	gm.AddEdge(edg1)
	gm.AddEdge(edg2)

	gm.Attach(hgw, msg, edg)
	gm.Attach(hgw, msg1, edg1)
	gm.Attach(msg2, hgw, edg2)

	neighbors, err := gm.GetNeighbors(hgw, socialGraph.SocialEdgeConst, socialGraph.SocialNodeConst)

	if err != nil {
		t.Error(err.Error())
	}

	if len(neighbors) != 3 {
		t.Errorf("GetNeighbors Failed - %v neighbors", len(neighbors))
	}
}
