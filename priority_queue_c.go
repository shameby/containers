package containers

import (
	"encoding/json"
)

type concurrencyPriorityQueue struct {
	l      []*Elem
	maxLen int
	RWLocker
}

func (npq concurrencyPriorityQueue) Top() *Elem {
	npq.RLock()
	defer npq.RUnlock()
	if npq.isEmpty() {
		return nil
	}
	return npq.l[0]
}

func (npq *concurrencyPriorityQueue) Push(ie IElem) bool {
	npq.Lock()
	defer npq.Unlock()
	if npq.isFull() {
		return false
	}
	e := initE(ie)
	npq.l = append(npq.l, e)
	npq.up(len(npq.l) - 1)
	return true
}

func (npq *concurrencyPriorityQueue) Pop() *Elem {
	npq.Lock()
	defer npq.Unlock()
	if npq.isEmpty() {
		return nil
	}
	e := npq.Top()
	npq.l[0] = npq.l[npq.len()-1]
	npq.l = npq.l[:npq.len()-1]
	npq.down(0)
	return e
}

func (npq *concurrencyPriorityQueue) up(start int) {
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

func (npq *concurrencyPriorityQueue) down(start int) {
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

func (npq *concurrencyPriorityQueue) GetList() []*Elem {
	npq.RLock()
	defer npq.RUnlock()
	return npq.l
}

func (npq concurrencyPriorityQueue) IsEmpty() bool {
	npq.RLock()
	defer npq.RUnlock()
	return npq.isEmpty()
}

func (npq concurrencyPriorityQueue) isEmpty() bool {
	if npq.len() == 0 {
		return true
	}
	return false
}

func (npq concurrencyPriorityQueue) IsFull() bool {
	npq.RLock()
	defer npq.RUnlock()
	return npq.isFull()
}

func (npq concurrencyPriorityQueue) Json() string {
	npq.RLock()
	defer npq.RUnlock()
	str, _ := json.Marshal(npq.l)
	return string(str)
}

func (npq concurrencyPriorityQueue) isFull() bool {
	if npq.maxLen != -1 && npq.len() == npq.maxLen {
		return true
	}
	return false
}

func (npq concurrencyPriorityQueue) Len() int {
	npq.RLock()
	defer npq.RUnlock()
	return npq.len()
}

func (npq concurrencyPriorityQueue) len() int {
	return len(npq.l)
}
