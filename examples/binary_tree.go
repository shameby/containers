package main

import (
	"fmt"
	"sync"

	c "../../containers"
)

type StuB struct {
	Name  string
	Grade string
	Score int64
}

func (s StuB) KV() (string, int64) {
	return s.Grade + ":" + s.Name, s.Score
}

func main() {
	students := []*StuB{
		{"tom", "grade1", 53},
		{"peter", "grade2", 80},
		{"jack", "grade1", 91},
		{"bob", "grade3", 91},
		{"huge", "grade3", 65},
		{"tim", "grade2", 77},
		{"randy", "grade2", 100},
		{"kya", "grade2", 99},
	}
	pq := c.NewBinaryTree(new(sync.RWMutex))
	wg := sync.WaitGroup{}
	for i := range students {
		wg.Add(1)
		go func(index int) {
			pq.Insert(students[index])
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println(pq.InorderTraversal())
	fmt.Println(pq.Search(91))
	fmt.Println(pq.Search(77))
	pq.Delete(80)
	fmt.Println(pq.InorderTraversal())
}
