package containers

type concurrencyPriorityQueue struct {
	n *normalPriorityQueue
	RWLocker
}

func (npq *concurrencyPriorityQueue) Top() *Elem {
	npq.RLock()
	defer npq.RUnlock()
	return npq.n.Top()
}

func (npq *concurrencyPriorityQueue) Push(ie IElem) bool {
	npq.Lock()
	defer npq.Unlock()
	return npq.n.Push(ie)
}

func (npq *concurrencyPriorityQueue) Pop() *Elem {
	npq.Lock()
	defer npq.Unlock()
	return npq.n.Pop()
}

func (npq *concurrencyPriorityQueue) GetList() []*Elem {
	npq.RLock()
	defer npq.RUnlock()
	return npq.n.GetList()
}

func (npq *concurrencyPriorityQueue) IsEmpty() bool {
	npq.RLock()
	defer npq.RUnlock()
	return npq.n.IsEmpty()
}

func (npq concurrencyPriorityQueue) IsFull() bool {
	npq.RLock()
	defer npq.RUnlock()
	return npq.n.IsFull()
}

func (npq concurrencyPriorityQueue) Json() string {
	npq.RLock()
	defer npq.RUnlock()
	return npq.n.Json()
}

func (npq concurrencyPriorityQueue) Len() int {
	npq.RLock()
	defer npq.RUnlock()
	return npq.n.Len()
}
