package containers

import (
	"unsafe"
)

// 队列
func NewQueue(len int, locker RWLocker) Queue {
	q := &queue{maxLen: len}
	if locker != nil {
		return &concurrencyQueue{n: q, RWLocker: locker}
	}
	return q
}

// lockFree 队列
func NewLockFreeQueue(maxLen int32) Queue {
	n := unsafe.Pointer(&node{})
	return &queueLF{head: n, tail: n, maxLen: maxLen}
}

// 集合
func NewSet(locker RWLocker) Set {
	s := &set{m: make(map[interface{}]int)}
	if locker != nil {
		return &concurrencySet{
			n:        s,
			RWLocker: locker,
		}
	}
	return s
}

// 栈
func NewStack(length int, locker RWLocker) Stack {
	s := &stack{l: make([]interface{}, length+1),
		topPtr: 0,
		maxLen: length,
	}
	if locker != nil {
		return &concurrencyStack{
			n:        s,
			RWLocker: locker,
		}
	}
	return s
}

// 底层为数组的双端队列
func NewDequeArr(k int, locker RWLocker) Deque {
	d := &dequeArr{
		maxLen: k + 1,
		data:   make([]interface{}, k+1), //空一个位置区分满和空
	}
	if locker != nil {
		return &concurrencyDequeArr{n: d, RWLocker: locker,}
	}
	return d
}

// 底层为链表的双端队列
func NewDequeL(maxLen int, locker RWLocker) Deque {
	d := &dequeL{curLen: 0, maxLen: maxLen}
	if locker != nil {
		return &concurrencyDequeL{
			n: d, RWLocker: locker,
		}
	}
	return d
}

// 二叉树
func NewBinaryTree(locker RWLocker) BinaryTree {
	b := &binaryTree{}
	if locker != nil {
		return &concurrencyBinaryTree{n: b, RWLocker: locker}
	}
	return b
}

// 底层为二叉堆的优先数列
func NewPriorityQueue(maxLen int, t HeapType, locker RWLocker) PriorityQueue {
	p := &normalPriorityQueue{t: t, maxLen: maxLen}
	if locker != nil {
		return &concurrencyPriorityQueue{n: p, RWLocker: locker}
	}
	return p
}
