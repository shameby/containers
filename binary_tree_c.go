package containers

type concurrencyBinaryTree struct {
	n *binaryTree
	RWLocker
}

func (cbt *concurrencyBinaryTree) Insert(elem IElem) BinaryTree {
	cbt.Lock()
	defer cbt.Unlock()
	return cbt.n.Insert(elem)
}

func (cbt *concurrencyBinaryTree) Search(i int64) *Elem {
	cbt.RLock()
	defer cbt.RUnlock()
	return cbt.n.Search(i)
}

func (cbt *concurrencyBinaryTree) Depth() int {
	cbt.RLock()
	defer cbt.RUnlock()
	return cbt.n.Depth()
}

func (cbt *concurrencyBinaryTree) Len() int {
	cbt.RLock()
	defer cbt.RUnlock()
	return cbt.n.Len()
}

func (cbt *concurrencyBinaryTree) InorderTraversal() (res []Elem) {
	cbt.RLock()
	defer cbt.RUnlock()
	return cbt.n.InorderTraversal()
}

func (cbt *concurrencyBinaryTree) PreorderTraversal() (res []Elem) {
	cbt.RLock()
	defer cbt.RUnlock()
	return cbt.n.PreorderTraversal()
}

func (cbt *concurrencyBinaryTree) PostorderTraversal() (res []Elem) {
	cbt.RLock()
	defer cbt.RUnlock()
	return cbt.n.PostorderTraversal()
}

func (cbt *concurrencyBinaryTree) Delete(i int64) int {
	cbt.Lock()
	defer cbt.Unlock()
	return cbt.n.Delete(i)
}
