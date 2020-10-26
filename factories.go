package data_structures

// 集合
func NewSet(isConcurrency bool) Set {
	if isConcurrency {
		return &concurrencySet{
			m: make(map[interface{}]int),
		}
	}
	return &set{
		m: make(map[interface{}]int),
	}
}

// 栈
func NewStack(length int, isConcurrency bool) Stack {
	if isConcurrency {
		return &concurrencyStack{
			l:      make([]interface{}, length+1),
			topPtr: 0,
			maxLen: length,
		}
	}
	return &stack{
		l:      make([]interface{}, length+1),
		topPtr: 0,
		maxLen: length,
	}
}

// 底层为数组的双端队列
func NewDequeArr(k int, isConcurrency bool) Deque {
	if isConcurrency {
		return &concurrencyDequeArr{
			length: k + 1,
			data:   make([]interface{}, k+1),
		}
	}
	return &dequeArr{
		length: k + 1,
		data:   make([]interface{}, k+1), //空一个位置区分满和空
	}
}

// 底层为链表的双端队列
func NewDequeL(k int, isConcurrency bool) Deque {
	if isConcurrency {
		return &concurrencyDequeL{
			curLen: 0, maxLen: k,
		}
	}
	return &dequeL{
		curLen: 0,
		maxLen: k,
	}
}

// 二叉树
func NewBinaryTree(isConcurrency bool) BinaryTree {
	if isConcurrency {
		return &concurrencyBinaryTree{}
	}
	return &binaryTree{}
}
