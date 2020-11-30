package containers

import "fmt"

type unionFind struct {
	cnt    int
	parent map[string]string
}

func (uf *unionFind) Append(s ...string) {
	for _, str := range s {
		if _, exist := uf.parent[str]; !exist {
			uf.parent[str] = str
			uf.cnt++
		}
	}
}

func (uf *unionFind) Union(a, b string) {
	rootA, rootB := uf.Find(a), uf.Find(b)
	fmt.Println(uf)
	if rootA != rootB {
		uf.parent[rootB] = rootA
		uf.cnt--
	}
}

func (uf *unionFind) Find(s string) string {
	cur := s
	for uf.parent[cur] != cur {
		cur = uf.parent[cur]
	}
	for uf.parent[s] != cur {
		s, uf.parent[s] = uf.parent[s], cur
	}
	return cur
}

func (uf unionFind) Count() int {
	return uf.cnt
}
