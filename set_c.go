package containers

type concurrencySet struct {
	n *set
	RWLocker
}

// concurrencySet
func (cs *concurrencySet) Add(i interface{}) bool {
	cs.Lock()
	defer cs.Unlock()
	return cs.n.Add(i)
}

func (cs *concurrencySet) Adds(is ...interface{}) bool {
	cs.Lock()
	defer cs.Unlock()
	return cs.n.Adds(is)
}

func (cs *concurrencySet) Delete(i interface{}) bool {
	cs.Lock()
	defer cs.Unlock()
	return cs.n.Delete(i)
}

func (cs concurrencySet) IsExist(i interface{}) bool {
	cs.RLock()
	defer cs.RUnlock()
	return cs.n.IsExist(i)
}

func (cs concurrencySet) ToList() IList {
	cs.RLock()
	defer cs.RUnlock()
	return cs.n.ToList()
}
