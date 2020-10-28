package containers

type Queue interface {
	Push(interface{}) bool
	Pop() interface{}
	Len() int
	IsEmpty() bool
	IsFull() bool
}

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

type BinaryTree interface {
	Insert(int64) BinaryTree
	Inserts(l []int64) BinaryTree
	Search(int64) int
	Delete(int64) int
	Depth() int
	Len() int
	InorderTraversal() []int64
	PreorderTraversal() []int64
	PostorderTraversal() []int64
}

type RWLocker interface {
	Lock()
	Unlock()
	RLock()
	RUnlock()
}