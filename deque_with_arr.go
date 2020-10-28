package containers

type dequeArr struct {
	maxLen int
	data   []interface{}
	head   int
	rear   int
}

func (dqr *dequeArr) InsertFront(value interface{}) bool {
	if dqr.isFull() {
		return false
	}
	if dqr.isEmpty() {
		if dqr.rear == dqr.maxLen-1 {
			dqr.rear = 0
		} else {
			dqr.rear++
		}
		dqr.data[dqr.head] = value
		return true
	}

	if dqr.head == 0 {
		dqr.head = dqr.maxLen - 1
	} else {
		dqr.head--
	}
	dqr.data[dqr.head] = value
	return true
}

func (dqr *dequeArr) InsertLast(value interface{}) bool {
	if dqr.isFull() {
		return false
	}
	if dqr.isEmpty() {
		dqr.data[dqr.rear] = value
		if dqr.rear == dqr.maxLen-1 {
			dqr.rear = 0
		} else {
			dqr.rear++
		}
		return true
	}

	dqr.data[dqr.rear] = value
	if dqr.rear == dqr.maxLen-1 {
		dqr.rear = 0
	} else {
		dqr.rear++
	}
	return true
}

func (dqr *dequeArr) DeleteFront() bool {
	if dqr.isEmpty() {
		return false
	}
	if dqr.head == dqr.maxLen-1 {
		dqr.head = 0
	} else {
		dqr.head++
	}
	return true
}

func (dqr *dequeArr) DeleteLast() bool {
	if dqr.isEmpty() {
		return false
	}
	if dqr.rear == 0 {
		dqr.rear = dqr.maxLen - 1
	} else {
		dqr.rear--
	}
	return true
}

func (dqr *dequeArr) GetFront() interface{} {
	if dqr.isEmpty() {
		return -1
	}
	return dqr.data[dqr.head]

}

func (dqr *dequeArr) GetRear() interface{} {
	if dqr.isEmpty() {
		return -1
	}
	if dqr.rear == 0 {
		return dqr.data[dqr.maxLen-1]
	}
	return dqr.data[dqr.rear-1]
}

func (dqr dequeArr) IsEmpty() bool {
	return dqr.isEmpty()
}

func (dqr dequeArr) isEmpty() bool {
	return dqr.head == dqr.rear
}

func (dqr dequeArr) IsFull() bool {
	return dqr.isFull()
}

func (dqr dequeArr) isFull() bool {
	return (dqr.rear+1)%dqr.maxLen == dqr.head
}

func (dqr dequeArr) ToList() IList {
	return dqr.data
}
