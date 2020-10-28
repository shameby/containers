package containers

type concurrencyStack struct {
	l      []interface{}
	topPtr int
	maxLen int
	RWLocker
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
	return cs.topPtr
}

func (cs concurrencyStack) ToList() IList {
	cs.RLock()
	defer cs.RUnlock()
	return cs.l
}

func (cs concurrencyStack) Top() interface{} {
	cs.RLock()
	defer cs.RUnlock()
	if cs.topPtr == 0 {
		return nil
	}
	return cs.l[cs.topPtr-1]
}
