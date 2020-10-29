package main

import (
	"fmt"
	"sync"

	c "../../containers"
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
	wg := sync.WaitGroup{}
	for i := range students {
		wg.Add(1)
		go func(index int) {
			pq.Push(students[index])
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println(pq.Json())
	fmt.Println("-----------------------")
	pq.Pop()
	fmt.Println(pq.Json())
	pq.Pop()
	fmt.Println(pq.Json())
	pq.Pop()
	fmt.Println(pq.Json())
}
