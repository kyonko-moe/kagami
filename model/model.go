package model

type NodeRegister interface {
	NodeByName(n string) *Node
	NodeByID(id int64) *Node
	SetNode(name string, id int64) error
	Clear()
}

type NodePicker interface {
	PickOne([]*Node) *Node
	PickList([]*Node) []*Node
}

type NodeNamer interface {
	ID(*Node) int64
	Region(*Node) int64
}
