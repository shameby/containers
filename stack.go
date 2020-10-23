package data_structures

import "sync"

type normalStack struct {
	l      []interface{}
	topPtr int
	maxLen int
}

type concurrencyStack struct {
	l      []interface{}
	topPtr int
	maxLen int
	sync.RWMutex
}

func NewStack(length int, isConcurrency bool) Stack {
	if isConcurrency {
		return &concurrencyStack{
			l:      make([]interface{}, length+1),
			topPtr: 0,
			maxLen: length,
		}
	}
	return &normalStack{
		l:      make([]interface{}, length+1),
		topPtr: 0,
		maxLen: length,
	}
}

// normalStack
func (ns *normalStack) Push(i interface{}) bool {
	if ns.topPtr == ns.maxLen {
		return false
	}

	ns.l[ns.topPtr] = i
	ns.topPtr++

	return true
}

func (ns *normalStack) Pop() (bool, interface{}) {
	if ns.topPtr == 0 {
		return false, nil
	}

	var res interface{}
	res, ns.l[ns.topPtr-1] = ns.l[ns.topPtr-1], 0
	ns.topPtr--

	return true, res
}

func (ns normalStack) Len() int {
	return len(ns.l)
}

func (ns normalStack) ToList() IList {
	return ns.l
}

func (ns normalStack) Top() interface{} {
	return ns.l[ns.topPtr-1]
}

// concurrencyStack
func (cs *concurrencyStack) Push(i interface{}) bool {
	cs.Lock()
	defer cs.Unlock()
	if cs.topPtr == cs.maxLen {
		return false
	}

	cs.l[cs.topPtr] = i
	cs.topPtr++

	return true
}

func (cs *concurrencyStack) Pop() (bool, interface{}) {
	cs.Lock()
	defer cs.Unlock()
	if cs.topPtr == 0 {
		return false, nil
	}

	var res interface{}
	res, cs.l[cs.topPtr-1] = cs.l[cs.topPtr-1], 0
	cs.topPtr--

	return true, res
}

func (cs concurrencyStack) Len() int {
	cs.RLock()
	defer cs.RUnlock()
	return len(cs.l)
}

func (cs concurrencyStack) ToList() IList {
	cs.RLock()
	defer cs.RUnlock()
	return cs.l
}

func (cs concurrencyStack) Top() interface{} {
	cs.RLock()
	defer cs.RUnlock()
	return cs.l[cs.topPtr-1]
}
