package containers

import (
	"encoding/json"
)

type normalPriorityQueue struct {
	l      []*Elem
	maxLen int
}

func (npq normalPriorityQueue) Top() *Elem {
	if npq.isEmpty() {
		return nil
	}
	return npq.l[0]
}

func (npq *normalPriorityQueue) Push(ie IElem) bool {
	if npq.isFull() {
		return false
	}
	e := initE(ie)
	npq.l = append(npq.l, e)
	npq.up(len(npq.l) - 1)
	return true
}

func (npq *normalPriorityQueue) Pop() *Elem {
	if npq.isEmpty() {
		return nil
	}
	e := npq.Top()
	npq.l[0] = npq.l[npq.len()-1]
	npq.l = npq.l[:npq.len()-1]
	npq.down(0)
	return e
}

func (npq *normalPriorityQueue) up(start int) {
	tmp := npq.l[start]
	cur := start
	par := (cur - 1) / 2
	for cur > 0 {
		if tmp.Score <= npq.l[par].Score {
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
	for child < npq.len() {
		if child < npq.len()-1 && npq.l[child].Score < npq.l[child+1].Score {
			child++
		}
		if tmp.Score > npq.l[child].Score {
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
	return npq.isEmpty()
}

func (npq normalPriorityQueue) isEmpty() bool {
	if npq.len() == 0 {
		return true
	}
	return false
}

func (npq normalPriorityQueue) IsFull() bool {
	return npq.isFull()
}

func (npq normalPriorityQueue) Json() string {
	str, _ := json.Marshal(npq.l)
	return string(str)
}

func (npq normalPriorityQueue) isFull() bool {
	if npq.maxLen != -1 && npq.len() == npq.maxLen {
		return true
	}
	return false
}

func (npq normalPriorityQueue) Len() int {
	return npq.len()
}

func (npq normalPriorityQueue) len() int {
	return len(npq.l)
}
