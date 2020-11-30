package containers

type concurrencyQueue struct {
	n *queue
	RWLocker
}

func (cq *concurrencyQueue) Push(i interface{}) bool {
	cq.Lock()
	defer cq.Unlock()
	return cq.n.Push(i)
}

func (cq *concurrencyQueue) Pop() interface{} {
	cq.Lock()
	defer cq.Unlock()
	return cq.n.Pop()
}

func (cq *concurrencyQueue) Len() int {
	cq.RLock()
	defer cq.RUnlock()
	return cq.n.Len()
}

func (cq *concurrencyQueue) IsFull() bool {
	cq.RLock()
	defer cq.RUnlock()
	return cq.n.IsFull()
}

func (cq *concurrencyQueue) IsEmpty() bool {
	cq.RLock()
	defer cq.RUnlock()
	return cq.n.IsEmpty()
}
