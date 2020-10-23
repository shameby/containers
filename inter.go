package data_structures

type Stack interface {
	Push(interface{}) bool
	Pop() (bool, interface{})
	Len() int
	ToList() IList
	Top() interface{}
}

type Set interface {
	Add(interface{}) bool
	Adds(...interface{}) bool
	Delete(interface{}) bool
	IsExist(interface{}) bool
	ToList() IList
}

type Deque interface {
	InsertFront(interface{}) bool
	InsertLast(interface{}) bool
	DeleteFront() bool
	DeleteLast() bool
	GetFront() interface{}
	GetRear() interface{}
	IsEmpty() bool
	IsFull() bool
	ToList() IList
}