package data_structures

import (
	"sync"
)

// 非线程安全的 Set, 无锁, 开销较小
type normalSet struct {
	m map[interface{}]int
}

// 线程安全的 Set, 自带读写锁, 开销较大
type concurrencySet struct {
	m map[interface{}]int
	sync.RWMutex
}

func NewSet(isConcurrency bool) Set {
	if isConcurrency {
		return &concurrencySet{
			m: make(map[interface{}]int),
		}
	}

	return &normalSet{
		m: make(map[interface{}]int),
	}

}

// normalSet
func (ns *normalSet) Add(i interface{}) bool {
	ns.m[i] = 0

	return true
}

func (ns *normalSet) Adds(is ...interface{}) bool {
	for _, i := range is {
		ns.m[i] = 0
	}

	return true
}

func (ns *normalSet) Delete(i interface{}) bool {
	if _, exist := ns.m[i]; !exist {
		return false
	}
	delete(ns.m, i)

	return true
}

func (ns normalSet) IsExist(i interface{}) bool {
	if _, exist := ns.m[i]; !exist {
		return false
	}

	return true
}

func (ns normalSet) ToList() IList {
	l := make([]interface{}, len(ns.m))
	for k := range ns.m {
		l = append(l, k)
	}

	return IList(l)
}

// concurrencySet
func (cs *concurrencySet) Add(i interface{}) bool {
	cs.Lock()
	defer cs.Unlock()
	cs.m[i] = 0

	return true
}

func (cs *concurrencySet) Adds(is ...interface{}) bool {
	cs.Lock()
	defer cs.Unlock()
	for _, i := range is {
		cs.m[i] = 0
	}

	return true
}

func (cs *concurrencySet) Delete(i interface{}) bool {
	cs.Lock()
	defer cs.Unlock()
	if _, exist := cs.m[i]; !exist {
		return false
	}
	delete(cs.m, i)

	return true
}

func (cs concurrencySet) IsExist(i interface{}) bool {
	cs.RLock()
	defer cs.RUnlock()
	if _, exist := cs.m[i]; !exist {
		return false
	}

	return true
}

func (cs concurrencySet) ToList() IList {
	cs.RLock()
	defer cs.RUnlock()
	l := make([]interface{}, len(cs.m))
	for k := range cs.m {
		l = append(l, k)
	}

	return IList(l)
}
