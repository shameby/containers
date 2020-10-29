package main

import (
	"fmt"

	c "../../containers"
	"sync"
	"time"
)

type Stu struct {
	Name  string
	Grade string
	Score int64
}

func (s Stu) KV() (string, int64) {
	return s.Grade + ":" + s.Name, s.Score
}

func main() {
	students := []Stu{
		{"tom", "grade1", 53},
		{"peter", "grade2", 80},
		{"jack", "grade1", 91},
		{"huge", "grade3", 65},
		{"tim", "grade2", 77},
		{"randy", "grade2", 100},
		{"kya", "grade2", 99},
	}
	pq := c.NewPriorityQueue(50, c.MaxRootHeap, new(sync.RWMutex))
	for _, stu := range students {
		go pq.Push(stu)
	}
	time.Sleep(500 * time.Millisecond)
	fmt.Println(pq.Json())
	fmt.Println("-----------------------")
	pq.Pop()
	fmt.Println(pq.Json())
	pq.Pop()
	fmt.Println(pq.Json())
	pq.Pop()
	fmt.Println(pq.Json())
}
