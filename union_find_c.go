package containers

type concurrencyUnionFind struct {
	n *unionFind
	RWLocker
}

func (cuf *concurrencyUnionFind) Append(s ...string) {
	cuf.Lock()
	defer cuf.Unlock()
	cuf.n.Append(s...)
}

func (cuf *concurrencyUnionFind) Union(a, b string) {
	cuf.Lock()
	defer cuf.Unlock()
	cuf.n.Union(a, b)
}

func (cuf *concurrencyUnionFind) Find(s string) string {
	cuf.Lock()
	defer cuf.Unlock()
	return cuf.n.Find(s)
}

func (cuf *concurrencyUnionFind) Count() int {
	cuf.RLock()
	defer cuf.RUnlock()
	return cuf.n.Count()
}
