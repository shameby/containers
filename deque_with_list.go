package containers

type dequeL struct {
	head   *duLNode
	tail   *duLNode
	curLen int
	maxLen int
}

func (dq *dequeL) InsertFront(value interface{}) bool {
	if dq.IsFull() {
		return false
	}
	dln := &duLNode{
		val: value,
	}
	if dq.IsEmpty() {
		dq.head, dq.tail = dln, dln
		dq.curLen++
		return true
	}
	dln.nex = dq.head
	dq.head.pre = dln
	dq.head = dln
	dq.curLen++
	return true
}

func (dq *dequeL) InsertLast(value interface{}) bool {
	if dq.IsFull() {
		return false
	}
	dln := &duLNode{
		val: value,
	}
	if dq.IsEmpty() {
		dq.head, dq.tail = dln, dln
		dq.curLen++
		return true
	}
	dln.pre = dq.tail
	dq.tail.nex = dln
	dq.tail = dln
	dq.curLen++
	return true
}

func (dq *dequeL) DeleteFront() bool {
	if dq.IsEmpty() {
		return false
	}
	if dq.head.nex == nil {
		dq.head, dq.tail = nil, nil
		dq.curLen--
		return true
	}
	lastHead := dq.head
	dq.head = lastHead.nex
	lastHead.nex = nil
	dq.head.pre = nil
	dq.curLen--
	return true
}

func (dq *dequeL) DeleteLast() bool {
	if dq.IsEmpty() {
		return false
	}
	if dq.tail.pre == nil {
		dq.head, dq.tail = nil, nil
		dq.curLen--
		return true
	}
	lastTail := dq.tail
	dq.tail = lastTail.pre
	lastTail.pre = nil
	dq.tail.nex = nil
	dq.curLen--
	return true
}

func (dq *dequeL) GetFront() interface{} {
	if dq.head == nil {
		return nil
	}
	return dq.head.val
}

func (dq *dequeL) GetRear() interface{} {
	if dq.tail == nil {
		return nil
	}
	return dq.tail.val
}

func (dq *dequeL) IsEmpty() bool {
	if dq.head == nil && dq.tail == nil {
		return true
	}
	return false
}

func (dq *dequeL) IsFull() bool {
	if dq.curLen == dq.maxLen {
		return true
	}
	return false
}

func (dq dequeL) ToList() IList {
	if dq.IsEmpty() {
		return nil
	}
	res := make([]interface{}, dq.curLen)
	cur := dq.head
	i := 0
	for cur != nil {
		res[i] = cur.val
		cur = cur.nex
		i++
	}
	return res
}