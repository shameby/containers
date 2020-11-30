package containers

type concurrencySkipList struct {
	RWLocker
	n *normalSkipList
}

func (cs *concurrencySkipList) Get(score float64) *SkElem {
	cs.RLock()
	defer cs.RUnlock()
	return cs.n.Get(score)
}

func (cs *concurrencySkipList) Set(ie IElem) bool {
	cs.Lock()
	defer cs.Unlock()
	return cs.n.Set(ie)
}

func (cs *concurrencySkipList) Delete(score float64) *SkElem {
	cs.Lock()
	defer cs.Unlock()
	return cs.n.Delete(score)
}

func (cs *concurrencySkipList) Len() int {
	cs.RLock()
	defer cs.RUnlock()
	return cs.n.Len()
}

func (cs *concurrencySkipList) Fmt() {
	cs.RLock()
	defer cs.RUnlock()
	cs.n.Fmt()
}

func (cs *concurrencySkipList) SetProbability(p float64) {
	cs.Lock()
	defer cs.Unlock()
	cs.n.SetProbability(p)
	return
}
