package containers

// 队列
func NewQueue(len int, locker RWLocker) Queue {
	if locker != nil {
		return &concurrencyQueue{maxLen: len, RWLocker: locker}
	}
	return &queue{maxLen: len}
}

// 集合
func NewSet(locker RWLocker) Set {
	if locker != nil {
		return &concurrencySet{
			m:        make(map[interface{}]int),
			RWLocker: locker,
		}
	}
	return &set{
		m: make(map[interface{}]int),
	}
}

// 栈
func NewStack(length int, locker RWLocker) Stack {
	if locker != nil {
		return &concurrencyStack{
			l:        make([]interface{}, length+1),
			topPtr:   0,
			maxLen:   length,
			RWLocker: locker,
		}
	}
	return &stack{
		l:      make([]interface{}, length+1),
		topPtr: 0,
		maxLen: length,
	}
}

// 底层为数组的双端队列
func NewDequeArr(k int, locker RWLocker) Deque {
	if locker != nil {
		return &concurrencyDequeArr{
			maxLen:   k + 1,
			data:     make([]interface{}, k+1),
			RWLocker: locker,
		}
	}
	return &dequeArr{
		maxLen: k + 1,
		data:   make([]interface{}, k+1), //空一个位置区分满和空
	}
}

// 底层为链表的双端队列
func NewDequeL(maxLen int, locker RWLocker) Deque {
	if locker != nil {
		return &concurrencyDequeL{
			curLen: 0, maxLen: maxLen, RWLocker: locker,
		}
	}
	return &dequeL{
		curLen: 0,
		maxLen: maxLen,
	}
}

// 二叉树
func NewBinaryTree(locker RWLocker) BinaryTree {
	if locker != nil {
		return &concurrencyBinaryTree{RWLocker: locker}
	}
	return &binaryTree{}
}

// 底层为二叉堆的优先数列
func NewPriorityQueue(maxLen int, locker RWLocker) PriorityQueue {
	if locker != nil {
		return &concurrencyPriorityQueue{RWLocker: locker}
	}
	return &normalPriorityQueue{maxLen: maxLen}
}
