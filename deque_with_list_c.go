package containers

type concurrencyDequeL struct {
	head   *duLNode
	tail   *duLNode
	curLen int
	maxLen int
	RWLocker
}

func (cdl *concurrencyDequeL) InsertFront(value interface{}) bool {
	cdl.Lock()
	defer cdl.Unlock()
	if cdl.isFull() {
		return false
	}
	dln := &duLNode{
		val: value,
	}
	if cdl.isEmpty() {
		cdl.head, cdl.tail = dln, dln
		cdl.curLen++
		return true
	}
	dln.nex = cdl.head
	cdl.head.pre = dln
	cdl.head = dln
	cdl.curLen++
	return true
}

func (cdl *concurrencyDequeL) InsertLast(value interface{}) bool {
	cdl.Lock()
	defer cdl.Unlock()
	if cdl.isFull() {
		return false
	}
	dln := &duLNode{
		val: value,
	}
	if cdl.isEmpty() {
		cdl.head, cdl.tail = dln, dln
		cdl.curLen++
		return true
	}
	dln.pre = cdl.tail
	cdl.tail.nex = dln
	cdl.tail = dln
	cdl.curLen++
	return true
}

func (cdl *concurrencyDequeL) DeleteFront() bool {
	cdl.Lock()
	defer cdl.Unlock()
	if cdl.isEmpty() {
		return false
	}

	if cdl.head.nex == nil {
		cdl.head, cdl.tail = nil, nil
		cdl.curLen--
		return true
	}

	lastHead := cdl.head
	cdl.head = lastHead.nex
	lastHead.nex = nil
	cdl.head.pre = nil
	cdl.curLen--

	return true
}

func (cdl *concurrencyDequeL) DeleteLast() bool {
	cdl.Lock()
	defer cdl.Unlock()
	if cdl.isEmpty() {
		return false
	}

	if cdl.tail.pre == nil {
		cdl.head, cdl.tail = nil, nil
		cdl.curLen--
		return true
	}

	lastTail := cdl.tail
	cdl.tail = lastTail.pre
	lastTail.pre = nil
	cdl.tail.nex = nil
	cdl.curLen--

	return true
}

func (cdl *concurrencyDequeL) GetFront() interface{} {
	cdl.RLock()
	defer cdl.RUnlock()
	if cdl.head == nil {
		return -1
	}
	return cdl.head.val
}

func (cdl *concurrencyDequeL) GetRear() interface{} {
	cdl.RLock()
	defer cdl.RUnlock()
	if cdl.tail == nil {
		return -1
	}
	return cdl.tail.val
}

func (cdl *concurrencyDequeL) IsEmpty() bool {
	cdl.RLock()
	defer cdl.RUnlock()
	return cdl.isEmpty()
}

func (cdl *concurrencyDequeL) isEmpty() bool {
	if cdl.head == nil && cdl.tail == nil {
		return true
	}
	return false
}

func (cdl *concurrencyDequeL) IsFull() bool {
	cdl.RLock()
	defer cdl.RUnlock()
	return cdl.isFull()
}

func (cdl *concurrencyDequeL) isFull() bool {
	if cdl.curLen == cdl.maxLen {
		return true
	}
	return false
}

func (cdl concurrencyDequeL) ToList() IList {
	cdl.RLock()
	defer cdl.RUnlock()
	if cdl.isEmpty() {
		return nil
	}
	res := make([]interface{}, cdl.curLen)
	cur := cdl.head
	i := 0
	for cur != nil {
		res[i] = cur.val
		cur = cur.nex
		i++
	}

	return res
}
