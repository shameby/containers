package containers

type concurrencyStack struct {
	n *stack
	RWLocker
}

// concurrencyStack
func (cs *concurrencyStack) Push(i interface{}) bool {
	cs.Lock()
	defer cs.Unlock()
	return cs.n.Push(i)
}

func (cs *concurrencyStack) Pop() (bool, interface{}) {
	cs.Lock()
	defer cs.Unlock()
	return cs.n.Pop()
}

func (cs concurrencyStack) Len() int {
	cs.RLock()
	defer cs.RUnlock()
	return cs.n.Len()
}

func (cs concurrencyStack) ToList() IList {
	cs.RLock()
	defer cs.RUnlock()
	return cs.n.ToList()
}

func (cs concurrencyStack) Top() interface{} {
	cs.RLock()
	defer cs.RUnlock()
	return cs.n.Top()
}
