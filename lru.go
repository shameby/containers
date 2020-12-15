package containers

import (
	"time"
	"fmt"
)

const (
	defaultTTL = 1 * time.Second
)

type linkNode struct {
	key, val  string
	unix      int64
	pre, next *linkNode
}

type lruCache struct {
	m          map[string]*linkNode
	cap        int
	ttl        time.Duration
	head, tail *linkNode
}

func (lc *lruCache) delete(node *linkNode) *lruCache {
	pre, next := node.pre, node.next
	pre.next, next.pre = next, pre
	return lc
}

func (lc *lruCache) add(node *linkNode) *lruCache {
	nxt := lc.head.next
	lc.head.next, node.next = node, nxt
	node.pre, nxt.pre = lc.head, node
	return lc
}

func (lc *lruCache) moveToHead(node *linkNode) *lruCache {
	lc.delete(node).add(node)
	return lc
}

func (lc *lruCache) Get(key string) string {
	if _, exist := lc.m[key]; !exist {
		return ""
	}
	lc.m[key].unix = time.Now().Add(lc.ttl).Unix()
	lc.moveToHead(lc.m[key])
	return lc.m[key].val
}

func (lc *lruCache) Delete(key string) {
	if _, exist := lc.m[key]; !exist {
		return
	}
	lc.delete(lc.m[key])
	delete(lc.m, key)
}

func (lc *lruCache) Put(key string, value string) {
	if _, exist := lc.m[key]; !exist {
		if len(lc.m) == lc.cap {
			delete(lc.m, lc.tail.pre.key)
			lc.delete(lc.tail.pre)
		}
		node := &linkNode{key, value, time.Now().Add(lc.ttl).Unix(), nil, nil}
		lc.m[key] = node
		lc.add(node)
	} else {
		lc.moveToHead(lc.m[key])
		lc.m[key].val = value
	}
}

func (lc *lruCache) TTL() {
	curr := lc.tail.pre
	for curr != lc.head {
		fmt.Println(*curr)
		pre := curr.pre
		if curr.unix != -1 && time.Now().Unix() > curr.unix {
			delete(lc.m, curr.key)
			lc.delete(curr)
			curr = pre
			continue
		}
		break
	}
}
