package containers

type concurrencyDequeL struct {
	n *dequeL
	RWLocker
}

func (cdl *concurrencyDequeL) InsertFront(value interface{}) bool {
	cdl.Lock()
	defer cdl.Unlock()
	return cdl.n.InsertFront(value)
}

func (cdl *concurrencyDequeL) InsertLast(value interface{}) bool {
	cdl.Lock()
	defer cdl.Unlock()
	return cdl.n.InsertLast(value)
}

func (cdl *concurrencyDequeL) DeleteFront() bool {
	cdl.Lock()
	defer cdl.Unlock()
	return cdl.n.DeleteFront()
}

func (cdl *concurrencyDequeL) DeleteLast() bool {
	cdl.Lock()
	defer cdl.Unlock()
	return cdl.n.DeleteLast()
}

func (cdl *concurrencyDequeL) GetFront() interface{} {
	cdl.RLock()
	defer cdl.RUnlock()
	return cdl.n.GetFront()
}

func (cdl *concurrencyDequeL) GetRear() interface{} {
	cdl.RLock()
	defer cdl.RUnlock()
	return cdl.n.GetRear()
}

func (cdl *concurrencyDequeL) IsEmpty() bool {
	cdl.RLock()
	defer cdl.RUnlock()
	return cdl.n.IsEmpty()
}

func (cdl *concurrencyDequeL) IsFull() bool {
	cdl.RLock()
	defer cdl.RUnlock()
	return cdl.n.IsFull()
}

func (cdl concurrencyDequeL) ToList() IList {
	cdl.RLock()
	defer cdl.RUnlock()
	return cdl.n.ToList()
}
