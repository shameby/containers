package data_structures

type lNode struct {
	val interface{}
	nex *lNode
}

type duLNode struct {
	val interface{}
	pre *duLNode
	nex *duLNode
}
