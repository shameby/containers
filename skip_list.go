package containers

import (
	"math"
	"math/rand"
	"fmt"
)

const DefaultProbability = 1 / math.E

type normalSkipList struct {
	maxLevel      int
	curLen        int
	randSource    rand.Source
	probability   float64
	proTable      []float64
	preNodesCache []*elementNode
	elementNode
}

func (ns normalSkipList) Get(score float64) *SkElem {
	var next *SkElem
	pre := &ns.elementNode
	for i := ns.maxLevel - 1; i >= 0; i-- {
		next = pre.next[i]
		for next != nil && score > next.Score {
			pre = &next.elementNode
			next = next.next[i]
		}
	}
	if next != nil && next.Score <= score {
		return next
	}
	return nil
}

func (ns *normalSkipList) Set(ie IElem) bool {
	ele := initSKE(ie, make([]*SkElem, ns.randLevel()))
	preList := ns.getPreElementNode(ele.Score)
	for i := range ele.next {
		ele.next[i] = preList[i].next[i]
		preList[i].next[i] = ele
	}
	ns.curLen++
	return false
}

func (ns *normalSkipList) Delete(score float64) *SkElem {
	preList := ns.getPreElementNode(score)
	if ele := preList[0].next[0]; ele != nil && ele.Score <= score {
		for k, v := range ele.next {
			preList[k].next[k] = v
		}
		ns.curLen--
		return ele
	}
	return nil
}

func (ns normalSkipList) Len() int {
	return ns.curLen
}

func (ns *normalSkipList) SetProbability(p float64) {
	ns.probability = p
	ns.proTable = probabilityTable(p, ns.maxLevel)
	return
}

func (ns normalSkipList) Fmt() {
	for i := ns.maxLevel - 1; i >= 0; i-- {
		cur := ns.next[i]
		for cur != nil {
			fmt.Print(cur.Key, cur.Score, " | ")
			cur = cur.next[i]
		}
		fmt.Println()
	}
}

func (ns normalSkipList) getPreElementNode(score float64) []*elementNode {
	pre, preList := &ns.elementNode, ns.preNodesCache
	for i := ns.maxLevel - 1; i >= 0; i-- {
		cur := pre.next[i]
		for cur != nil && score > cur.Score {
			pre = &cur.elementNode
			cur = cur.next[i]
		}
		preList[i] = pre
	}
	return preList
}

func (ns normalSkipList) randLevel() (level int) {
	r := float64(ns.randSource.Int63()) / (1 << 63)
	level = 1
	for level < ns.maxLevel && r < ns.proTable[level] {
		level++
	}
	return
}

func probabilityTable(probability float64, maxLevel int) (table []float64) {
	for i := 1; i <= maxLevel; i++ {
		prob := math.Pow(probability, float64(i-1))
		table = append(table, prob)
	}
	return table
}
