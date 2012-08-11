package spiderDB

type Edge interface {
	GetId() string
	SetId(id string)
	GetWeight() int
	SetWeight(weight int)
	GetType() string
	SetType(typestr string)
}
