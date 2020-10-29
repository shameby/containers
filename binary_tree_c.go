package containers

type concurrencyBinaryTree struct {
	n *binaryTree
	RWLocker
}

func (cbt *concurrencyBinaryTree) Insert(i int64) BinaryTree {
	cbt.Lock()
	defer cbt.Unlock()
	return cbt.n.Insert(i)
}

func (cbt *concurrencyBinaryTree) Inserts(l []int64) BinaryTree {
	cbt.Lock()
	defer cbt.Unlock()
	return cbt.n.Inserts(l)
}

func (cbt *concurrencyBinaryTree) Search(i int64) int {
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

func (cbt *concurrencyBinaryTree) InorderTraversal() (res []int64) {
	cbt.RLock()
	defer cbt.RUnlock()
	return cbt.n.InorderTraversal()
}

func (cbt *concurrencyBinaryTree) PreorderTraversal() (res []int64) {
	cbt.RLock()
	defer cbt.RUnlock()
	return cbt.n.PreorderTraversal()
}

func (cbt *concurrencyBinaryTree) PostorderTraversal() (res []int64) {
	cbt.RLock()
	defer cbt.RUnlock()
	return cbt.n.PostorderTraversal()
}

func (cbt *concurrencyBinaryTree) Delete(i int64) int {
	cbt.Lock()
	defer cbt.Unlock()
	return cbt.n.Delete(i)
}
