package containers

import (
	"encoding/json"
)

const (
	// 小根堆
	MinRootHeap HeapType = 1 + iota
	// 大根堆
	MaxRootHeap
)

type HeapType uint8

type normalPriorityQueue struct {
	l      []*Elem
	t      HeapType
	maxLen int
}

func (npq normalPriorityQueue) Top() *Elem {
	if npq.IsEmpty() {
		return nil
	}
	return npq.l[0]
}

func (npq *normalPriorityQueue) Push(ie IElem) bool {
	if npq.IsFull() {
		return false
	}
	e := initE(ie)
	npq.l = append(npq.l, e)
	npq.up(len(npq.l) - 1)
	return true
}

func (npq *normalPriorityQueue) Pop() *Elem {
	if npq.IsEmpty() {
		return nil
	}
	e := npq.Top()
	npq.l[0] = npq.l[npq.Len()-1]
	npq.l = npq.l[:npq.Len()-1]
	npq.down(0)
	return e
}

func (npq *normalPriorityQueue) up(start int) {
	tmp := npq.l[start]
	cur := start
	par := (cur - 1) / 2
	for cur > 0 {
		if npq.t == MaxRootHeap && tmp.Score <= npq.l[par].Score {
			break
		}
		if npq.t == MinRootHeap && tmp.Score >= npq.l[par].Score {
			break
		}
		npq.l[cur] = npq.l[par]
		cur = par
		par = (cur - 1) / 2
	}
	npq.l[cur] = tmp
}

func (npq *normalPriorityQueue) down(start int) {
	tmp := npq.l[start]
	cur := start
	child := 2*cur + 1
	for child < npq.Len() {
		if child < npq.Len()-1 {
			if npq.t == MaxRootHeap && npq.l[child].Score < npq.l[child+1].Score {
				child++
			}
			if npq.t == MinRootHeap && npq.l[child].Score > npq.l[child+1].Score {
				child++
			}
		}
		if npq.t == MaxRootHeap && tmp.Score >= npq.l[child].Score {
			break
		}
		if npq.t == MinRootHeap && tmp.Score <= npq.l[child].Score {
			break
		}
		npq.l[cur] = npq.l[child]
		cur = child
		child = 2*child + 1
	}
	npq.l[cur] = tmp
}

func (npq *normalPriorityQueue) GetList() []*Elem {
	return npq.l
}

func (npq normalPriorityQueue) IsEmpty() bool {
	if npq.Len() == 0 {
		return true
	}
	return false
}

func (npq normalPriorityQueue) IsFull() bool {
	if npq.maxLen != -1 && npq.Len() == npq.maxLen {
		return true
	}
	return false
}

func (npq normalPriorityQueue) Json() string {
	str, _ := json.Marshal(npq.l)
	return string(str)
}

func (npq normalPriorityQueue) Len() int {
	return len(npq.l)
}
