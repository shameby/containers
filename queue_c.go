package containers

type concurrencyQueue struct {
	head   *lNode
	tail   *lNode
	maxLen int
	curLen int
	RWLocker
}

func (cq *concurrencyQueue) Push(i interface{}) bool {
	cq.Lock()
	defer cq.Unlock()
	if cq.isFull() {
		return false
	}
	if i == nil {
		return false
	}
	ln := &lNode{val: i}
	if cq.head == nil {
		cq.head = ln
		cq.tail = cq.head
	} else {
		cq.tail.nex = ln
		cq.tail = ln
	}
	cq.curLen++
	return true
}

func (cq *concurrencyQueue) Pop() interface{} {
	cq.Lock()
	defer cq.Unlock()
	if cq.isEmpty() {
		return nil
	}
	node := cq.head
	cq.head = node.nex
	cq.curLen--
	return node.val
}

func (cq concurrencyQueue) Len() int {
	cq.RLock()
	defer cq.RUnlock()
	return cq.len()
}

func (cq concurrencyQueue) len() int {
	return cq.curLen
}

func (cq concurrencyQueue) IsFull() bool {
	cq.RLock()
	defer cq.RUnlock()
	return cq.isFull()
}

func (cq concurrencyQueue) isFull() bool {
	cq.RLock()
	defer cq.RUnlock()
	if cq.maxLen != -1 && cq.maxLen == cq.curLen {
		return true
	}
	return false
}

func (cq concurrencyQueue) IsEmpty() bool {
	cq.RLock()
	defer cq.RUnlock()
	return cq.isEmpty()
}

func (cq concurrencyQueue) isEmpty() bool {
	cq.RLock()
	defer cq.RUnlock()
	if cq.curLen == 0 {
		return true
	}
	return false
}
