package main

import (
	"fmt"
	c "containers"
	"time"
	"sync"
)

func main() {
	lc := c.NewLRU(10000, time.Second, new(sync.Mutex))
	lc.Put("abc", "abc's data")
	lc.Put("dfs", "dfs's data")
	lc.Put("hhh", "hhh's data")
	lc.Put("iii", "iii's data")
	fmt.Println(lc.Get("abc"))
	fmt.Println(lc.Get("fdd"))
	fmt.Println(lc.Get("hhh"))
	lc.Delete("abc")
	time.Sleep(2 * time.Second)
	lc.TTL()
}