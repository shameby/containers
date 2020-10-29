package containers

type queue struct {
	head   *lNode
	tail   *lNode
	maxLen int
	curLen int
}

func (q *queue) Push(i interface{}) bool {
	if q.IsFull() {
		return false
	}
	if i == nil {
		return false
	}
	ln := &lNode{val: i}
	if q.head == nil {
		q.head = ln
		q.tail = q.head
	} else {
		q.tail.nex = ln
		q.tail = ln
	}
	q.curLen++
	return true
}

func (q *queue) Pop() interface{} {
	if q.IsEmpty() {
		return nil
	}
	node := q.head
	q.head = node.nex
	q.curLen--
	return node.val
}

func (q queue) Len() int {
	return q.curLen
}

func (q queue) IsFull() bool {
	if q.maxLen != -1 && q.maxLen == q.curLen {
		return true
	}
	return false
}

func (q queue) IsEmpty() bool {
	if q.curLen == 0 {
		return true
	}
	return false
}
