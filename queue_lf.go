package containers

import (
	"unsafe"
	"sync/atomic"
)

type queueLF struct {
	head   unsafe.Pointer
	tail   unsafe.Pointer
	curLen int32
	maxLen int32
}

type node struct {
	value interface{}
	next  unsafe.Pointer
}

func (ql *queueLF) Push(i interface{}) bool {
	if ql.IsFull() {
		return false
	}
	n := &node{value: i}
	for {
		tail := load(&ql.tail)
		next := load(&tail.next)
		if tail == load(&ql.tail) {
			if next == nil {
				if cas(&tail.next, next, n) {
					cas(&ql.tail, tail, n)
					atomic.AddInt32(&ql.curLen, 1)
					return true
				}
			} else {
				cas(&ql.tail, tail, next)
			}
		}
	}
}

func (ql *queueLF) Pop() interface{} {
	for {
		head := load(&ql.head)
		next := load(&head.next)
		if head == load(&ql.head) {
			if ql.IsEmpty() && next == nil {
				return nil
			} else {
				v := next.value
				if cas(&ql.head, head, next) {
					atomic.AddInt32(&ql.curLen, -1)
					return v
				}
			}
		}
	}
}

func (ql *queueLF) Len() int {
	return int(atomic.LoadInt32(&ql.curLen))
}

func (ql *queueLF) IsFull() bool {
	return atomic.LoadInt32(&ql.curLen) == ql.maxLen
}

func (ql *queueLF) IsEmpty() bool {
	return load(&ql.head) == load(&ql.tail)
}

func load(p *unsafe.Pointer) (n *node) {
	return (*node)(atomic.LoadPointer(p))
}

// 封装CAS
func cas(p *unsafe.Pointer, old, new *node) (ok bool) {
	return atomic.CompareAndSwapPointer(p, unsafe.Pointer(old), unsafe.Pointer(new))
}
