package containers

type concurrencyDequeArr struct {
	n *dequeArr
	RWLocker
}

func (cdr *concurrencyDequeArr) InsertFront(value interface{}) bool {
	cdr.Lock()
	defer cdr.Unlock()
	return cdr.n.InsertFront(value)
}

func (cdr *concurrencyDequeArr) InsertLast(value interface{}) bool {
	cdr.Lock()
	defer cdr.Unlock()
	return cdr.n.InsertLast(value)
}

func (cdr *concurrencyDequeArr) DeleteFront() bool {
	cdr.Lock()
	defer cdr.Unlock()
	return cdr.n.DeleteFront()
}

func (cdr *concurrencyDequeArr) DeleteLast() bool {
	cdr.Lock()
	defer cdr.Unlock()
	return cdr.n.DeleteLast()
}

func (cdr *concurrencyDequeArr) GetFront() interface{} {
	cdr.RLock()
	defer cdr.RUnlock()
	return cdr.n.GetFront()
}

func (cdr *concurrencyDequeArr) GetRear() interface{} {
	cdr.RLock()
	defer cdr.RUnlock()
	return cdr.n.GetRear()
}

func (cdr *concurrencyDequeArr) IsEmpty() bool {
	cdr.RLock()
	defer cdr.RUnlock()
	return cdr.n.IsEmpty()
}

func (cdr *concurrencyDequeArr) IsFull() bool {
	cdr.RLock()
	defer cdr.RUnlock()
	return cdr.n.IsFull()
}

func (cdr *concurrencyDequeArr) ToList() IList {
	cdr.RLock()
	defer cdr.RUnlock()
	return cdr.n.ToList()
}
