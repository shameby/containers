package containers

import "sync"

type concurrencySet struct {
	m map[interface{}]int
	RWLocker
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
