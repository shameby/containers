package containers

type normalTrie struct {
	*trie
	cnt int
}

type trie struct {
	next  [24]*trie
	isEnd bool
}

func (nt *normalTrie) Insert(word string) {
	cur := nt.trie
	for _, v := range word {
		v = v - 'a'
		if cur.next[v] == nil {
			cur.next[v] = &trie{}
		}
		cur = cur.next[v]
	}
	cur.isEnd = true
	nt.cnt++
}

func (nt normalTrie) Search(word string) bool {
	cur := nt.trie
	for _, v := range word {
		v = v - 'a'
		if cur.next[v] == nil {
			return false
		}
		cur = cur.next[v]
	}
	return cur.isEnd
}

func (nt normalTrie) Len() int {
	return nt.cnt
}

func (nt normalTrie) StartWith(word string) bool {
	cur := nt.trie
	for _, v := range word {
		v = v - 'a'
		if cur.next[v] == nil {
			return false
		}
		cur = cur.next[v]
	}
	return true
}
