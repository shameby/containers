package containers

type concurrencyDequeArr struct {
	maxLen int
	data   []interface{}
	head   int
	rear   int
	RWLocker
}

func (cdr *concurrencyDequeArr) InsertFront(value interface{}) bool {
	cdr.Lock()
	defer cdr.Unlock()
	if cdr.isFull() {
		return false
	}
	if cdr.isEmpty() {
		if cdr.rear == cdr.maxLen-1 {
			cdr.rear = 0
		} else {
			cdr.rear++
		}
		cdr.data[cdr.head] = value
		return true
	}

	if cdr.head == 0 {
		cdr.head = cdr.maxLen - 1
	} else {
		cdr.head--
	}
	cdr.data[cdr.head] = value
	return true
}

func (cdr *concurrencyDequeArr) InsertLast(value interface{}) bool {
	cdr.Lock()
	defer cdr.Unlock()
	if cdr.isFull() {
		return false
	}
	if cdr.isEmpty() {
		cdr.data[cdr.rear] = value
		if cdr.rear == cdr.maxLen-1 {
			cdr.rear = 0
		} else {
			cdr.rear++
		}
		return true
	}

	cdr.data[cdr.rear] = value
	if cdr.rear == cdr.maxLen-1 {
		cdr.rear = 0
	} else {
		cdr.rear++
	}
	return true
}

func (cdr *concurrencyDequeArr) DeleteFront() bool {
	cdr.Lock()
	defer cdr.Unlock()
	if cdr.isEmpty() {
		return false
	}
	if cdr.head == cdr.maxLen-1 {
		cdr.head = 0
	} else {
		cdr.head++
	}
	return true
}

func (cdr *concurrencyDequeArr) DeleteLast() bool {
	cdr.Lock()
	defer cdr.Unlock()
	if cdr.isEmpty() {
		return false
	}
	if cdr.rear == 0 {
		cdr.rear = cdr.maxLen - 1
	} else {
		cdr.rear--
	}
	return true
}

func (cdr *concurrencyDequeArr) GetFront() interface{} {
	cdr.RLock()
	defer cdr.RUnlock()
	if cdr.isEmpty() {
		return -1
	}
	return cdr.data[cdr.head]

}

func (cdr *concurrencyDequeArr) GetRear() interface{} {
	cdr.RLock()
	defer cdr.RUnlock()
	if cdr.isEmpty() {
		return -1
	}
	if cdr.rear == 0 {
		return cdr.data[cdr.maxLen-1]
	}
	return cdr.data[cdr.rear-1]
}

func (cdr concurrencyDequeArr) IsEmpty() bool {
	cdr.RLock()
	defer cdr.RUnlock()
	return cdr.isEmpty()
}

func (cdr concurrencyDequeArr) isEmpty() bool {
	return cdr.head == cdr.rear
}

func (cdr concurrencyDequeArr) IsFull() bool {
	cdr.RLock()
	defer cdr.RUnlock()
	return cdr.isFull()
}

func (cdr concurrencyDequeArr) isFull() bool {
	return (cdr.rear+1)%cdr.maxLen == cdr.head
}

func (cdr concurrencyDequeArr) ToList() IList {
	cdr.RLock()
	defer cdr.RUnlock()
	return cdr.data
}
