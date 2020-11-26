package containers

// normal elem
type Elem struct {
	Key   string  `json:"key"`
	Score float64 `json:"score"`
}

func initE(ie IElem) *Elem {
	k, v := ie.KV()
	return &Elem{k, v}
}

// elem just for skipList
type SkElem struct {
	Key   string  `json:"key"`
	Score float64 `json:"score"`
	elementNode
}

type elementNode struct {
	next []*SkElem
}

func (se SkElem) Next() *SkElem {
	return se.next[0]
}

func initSKE(ie IElem, next []*SkElem) *SkElem {
	k, v := ie.KV()
	return &SkElem{k, v, elementNode{next}}
}
