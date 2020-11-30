package main

import (
	"sync"
	"fmt"
	c "../../containers"
)

func main() {
	uf := c.NewUnionFind(&sync.RWMutex{})
	uf.Append([]string{"apple", "pea", "tomato"}...)
	fmt.Println(uf.Count())
	uf.Union("apple", "pea")
	fmt.Println(uf.Count())
	uf.Union("apple", "pea")
	fmt.Println(uf.Count())
}