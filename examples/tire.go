package main

import (
	"fmt"
	c "../../containers"
	"sync"
)

func main() {
	t := c.NewTrie(&sync.RWMutex{})
	t.Insert("apple")
	t.Insert("pea")
	t.Insert("peach")
	fmt.Println(t.Len())
	fmt.Println(t.Search("apple"))
	fmt.Println(t.Search("app"))
	fmt.Println(t.StartWith("app"))
	fmt.Println(t.StartWith("pea"))
	fmt.Println(t.Search("pea"))
	fmt.Println(t.Search("peach"))
}
