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
	Insert(IElem) BinaryTree
	Search(float64) *Elem
	Delete(float64) int
	Depth() int
	Len() int
	InorderTraversal() []Elem
	PreorderTraversal() []Elem
	PostorderTraversal() []Elem
}

type PriorityQueue interface {
	Top() *Elem
	Push(IElem) bool
	Pop() *Elem
	IsEmpty() bool
	IsFull() bool
	Len() int
	GetList() []*Elem
	Json() string
}

type SkipList interface {
	Get(float64) *SkElem
	Set(IElem) bool
	Delete(float64) *SkElem
	Len() int
	SetProbability(float64)
}

type RWLocker interface {
	Lock()
	Unlock()
	RLock()
	RUnlock()
}

type IElem interface {
	KV() (string, float64)
}
