package data_structures

type dequeArr struct {
	length int
	data   []interface{}
	head   int
	rear   int
}

// 底层为数组的双端队列
func NewDequeArr(k int) Deque {
	return &dequeArr{
		length: k + 1,
		data:   make([]interface{}, k+1), //空一个位置区分满和空
		head:   0,
		rear:   0,
	}
}

func (dqr *dequeArr) InsertFront(value interface{}) bool {
	if dqr.IsFull() {
		return false
	}
	if dqr.IsEmpty() {
		if dqr.rear == dqr.length-1 {
			dqr.rear = 0
		} else {
			dqr.rear++
		}
		dqr.data[dqr.head] = value
		return true
	}

	if dqr.head == 0 {
		dqr.head = dqr.length - 1
	} else {
		dqr.head--
	}
	dqr.data[dqr.head] = value
	return true
}

func (dqr *dequeArr) InsertLast(value interface{}) bool {
	if dqr.IsFull() {
		return false
	}
	if dqr.IsEmpty() {
		dqr.data[dqr.rear] = value
		if dqr.rear == dqr.length-1 {
			dqr.rear = 0
		} else {
			dqr.rear++
		}
		return true
	}

	dqr.data[dqr.rear] = value
	if dqr.rear == dqr.length-1 {
		dqr.rear = 0
	} else {
		dqr.rear++
	}
	return true
}

func (dqr *dequeArr) DeleteFront() bool {
	if dqr.IsEmpty() {
		return false
	}
	if dqr.head == dqr.length-1 {
		dqr.head = 0
	} else {
		dqr.head++
	}
	return true
}

func (dqr *dequeArr) DeleteLast() bool {
	if dqr.IsEmpty() {
		return false
	}
	if dqr.rear == 0 {
		dqr.rear = dqr.length - 1
	} else {
		dqr.rear--
	}
	return true
}

func (dqr *dequeArr) GetFront() interface{} {
	if dqr.IsEmpty() {
		return -1
	}
	return dqr.data[dqr.head]

}

func (dqr *dequeArr) GetRear() interface{} {
	if dqr.IsEmpty() {
		return -1
	}
	if dqr.rear == 0 {
		return dqr.data[dqr.length-1]
	}
	return dqr.data[dqr.rear-1]
}

func (dqr dequeArr) IsEmpty() bool {
	return dqr.head == dqr.rear
}

func (dqr dequeArr) IsFull() bool {
	return (dqr.rear+1)%dqr.length == dqr.head
}

func (dqr dequeArr) ToList() IList {
	return dqr.data
}
