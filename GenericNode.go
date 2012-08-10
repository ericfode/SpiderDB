package spiderDB

type GenericNode struct {
	id    string
	value int
	props map[string]*interface{}
}

func (n *GenericNode) Update() {

}

func (n *GenericNode) GetPropMap() (map[string][]byte){

}

func (n *GenericNode) SetPropMap(map[string][]byte){

}

