package containers

type concurrencyTrie struct {
	n *normalTrie
	RWLocker
}

func (ct *concurrencyTrie) Insert(word string) {
	ct.Lock()
	defer ct.Unlock()
	ct.n.Insert(word)
}

func (ct *concurrencyTrie) Search(word string) bool {
	ct.RLock()
	defer ct.RUnlock()
	return ct.n.Search(word)
}

func (ct *concurrencyTrie) Len() int {
	ct.RLock()
	defer ct.RUnlock()
	return ct.n.Len()
}

func (ct concurrencyTrie) StartWith(word string) bool {
	ct.RLock()
	defer ct.RUnlock()
	return ct.n.StartWith(word)
}
