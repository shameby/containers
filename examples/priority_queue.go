package main

import (
	"fmt"

	c "../../containers"
)

type Stu struct {
	Name string
	Grade string
	Score int64
}

func (s Stu) KV() (string, int64) {
	return s.Grade + ":" + s.Name, s.Score
}

func main() {
	students := []*Stu{
		{"tom", "grade1", 53},
		{"peter", "grade2", 80},
		{"jack", "grade1", 91},
		{"huge", "grade3", 65},
		{"tim", "grade2", 77},
		{"randy", "grade2", 100},
		{"kya", "grade2", 99},
	}
	pq := c.NewPriorityQueue(50, nil)
	for _, stu := range students {
		pq.Push(stu)
		fmt.Println(pq.Json())
	}
	pq.Pop()
	fmt.Println(pq.Json())
	pq.Pop()
	fmt.Println(pq.Json())
	pq.Pop()
	fmt.Println(pq.Json())
}
