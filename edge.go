package spiderDB

type Edge interface{
	GetId() string
	SetId(id string)
	GetWeight() int
	SetWeight(int weight)
	GetType() string
	SetType(type string)
}

// GETTERS & SETTERS

//Return connected nodes
func (e *Edge) GetNodes() {

}

func (e *Edge) SetDesc() {

}

func (e *Edge) GetDesc() {

}

func (e *Edge) GetID() {

}

// Update local copy from database
func (e *Edge) Pull() {

}
