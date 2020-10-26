package data_structures

type concurrencyDequeArr struct {
	length int
	data   []interface{}
	head   int
	rear   int
}

func (cdr *concurrencyDequeArr) InsertFront(value interface{}) bool {
	if cdr.IsFull() {
		return false
	}
	if cdr.IsEmpty() {
		if cdr.rear == cdr.length-1 {
			cdr.rear = 0
		} else {
			cdr.rear++
		}
		cdr.data[cdr.head] = value
		return true
	}

	if cdr.head == 0 {
		cdr.head = cdr.length - 1
	} else {
		cdr.head--
	}
	cdr.data[cdr.head] = value
	return true
}

func (cdr *concurrencyDequeArr) InsertLast(value interface{}) bool {
	if cdr.IsFull() {
		return false
	}
	if cdr.IsEmpty() {
		cdr.data[cdr.rear] = value
		if cdr.rear == cdr.length-1 {
			cdr.rear = 0
		} else {
			cdr.rear++
		}
		return true
	}

	cdr.data[cdr.rear] = value
	if cdr.rear == cdr.length-1 {
		cdr.rear = 0
	} else {
		cdr.rear++
	}
	return true
}

func (cdr *concurrencyDequeArr) DeleteFront() bool {
	if cdr.IsEmpty() {
		return false
	}
	if cdr.head == cdr.length-1 {
		cdr.head = 0
	} else {
		cdr.head++
	}
	return true
}

func (cdr *concurrencyDequeArr) DeleteLast() bool {
	if cdr.IsEmpty() {
		return false
	}
	if cdr.rear == 0 {
		cdr.rear = cdr.length - 1
	} else {
		cdr.rear--
	}
	return true
}

func (cdr *concurrencyDequeArr) GetFront() interface{} {
	if cdr.IsEmpty() {
		return -1
	}
	return cdr.data[cdr.head]

}

func (cdr *concurrencyDequeArr) GetRear() interface{} {
	if cdr.IsEmpty() {
		return -1
	}
	if cdr.rear == 0 {
		return cdr.data[cdr.length-1]
	}
	return cdr.data[cdr.rear-1]
}

func (cdr concurrencyDequeArr) IsEmpty() bool {
	return cdr.head == cdr.rear
}

func (cdr concurrencyDequeArr) IsFull() bool {
	return (cdr.rear+1)%cdr.length == cdr.head
}

func (cdr concurrencyDequeArr) ToList() IList {
	return cdr.data
}
