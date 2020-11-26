package main

import (
	c "../../containers"
	"fmt"
	"sync"
)

type Teacher struct {
	School string
	Name   string
	Score  float64
}

func (t Teacher) KV() (string, float64) {
	return t.School+":"+t.Name, t.Score
}

func main() {
	tList := []*Teacher{
		{"一中", "王老师", 67.2},
		{"二中", "李老师", 70.1},
		{"三中", "张老师", 88.3},
		{"一中", "赵老师", 75.9},
		{"二中", "严老师", 91.4},
	}
	list := c.NewSkipList(4, new(sync.RWMutex))
	list.Set(tList[0])
	list.Set(tList[1])
	list.Set(tList[2])
	list.Set(tList[3])
	list.Set(tList[4])
	fmt.Println(list.Get(91.4))
}
