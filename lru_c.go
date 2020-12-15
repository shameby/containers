package containers

import (
	"sync"
)

type currencyLRU struct {
	n *lruCache
	sync.Locker
}

func (lc *currencyLRU) Get(key string) string {
	lc.Lock()
	defer lc.Unlock()
	return lc.n.Get(key)
}

func (lc *currencyLRU) Delete(key string) {
	lc.Lock()
	defer lc.Unlock()
	lc.n.Delete(key)
}

func (lc *currencyLRU) Put(key string, value string) {
	lc.Lock()
	defer lc.Unlock()
	lc.n.Put(key, value)
}

func (lc *currencyLRU) TTL() {
	lc.Lock()
	defer lc.Unlock()
	lc.n.TTL()
}
