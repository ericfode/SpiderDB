package spiderDB

type Edge interface{
	GetId() string
	SetId(id string)
	GetWeight() int
	SetWeight(int weight)
	GetType() string
	SetType(type string)
}

