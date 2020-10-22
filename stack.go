package data_structures

import "sync"

type Stack interface {
	Push(interface{}) bool
	Pop() (bool, interface{})
	Len() int
	ToList() IList
	Top() interface{}
}

type normalStack struct {
	l      []interface{}
	topPtr int
	curLen int
}

type concurrencyStack struct {
	l      []interface{}
	topPtr int
	curLen int
	sync.RWMutex
}

func NewStack(length int, isConcurrency bool) Stack {
	if isConcurrency {
		return &concurrencyStack{
			l:      make([]interface{}, length+1),
			topPtr: 0,
			curLen: length,
		}
	}
	return &normalStack{
		l:      make([]interface{}, length+1),
		topPtr: 0,
		curLen: length,
	}
}

// normalStack
func (ns *normalStack) Push(i interface{}) bool {
	if ns.topPtr == ns.curLen+1 {
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
func (ns *concurrencyStack) Push(i interface{}) bool {
	ns.Lock()
	defer ns.Unlock()
	if ns.topPtr == ns.curLen+1 {
		return false
	}

	ns.l[ns.topPtr] = i
	ns.topPtr++

	return true
}

func (ns *concurrencyStack) Pop() (bool, interface{}) {
	ns.Lock()
	defer ns.Unlock()
	if ns.topPtr == 0 {
		return false, nil
	}

	var res interface{}
	res, ns.l[ns.topPtr-1] = ns.l[ns.topPtr-1], 0
	ns.topPtr--

	return true, res
}

func (ns concurrencyStack) Len() int {
	ns.RLock()
	defer ns.RUnlock()
	return len(ns.l)
}

func (ns concurrencyStack) ToList() IList {
	ns.RLock()
	defer ns.RUnlock()
	return ns.l
}

func (ns concurrencyStack) Top() interface{} {
	ns.RLock()
	defer ns.RUnlock()
	return ns.l[ns.topPtr-1]
}
