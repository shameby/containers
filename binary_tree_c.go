package containers

type concurrencyBinaryTree struct {
	root   *binaryTreeNode
	curLen int
	RWLocker
}

func (cbt *concurrencyBinaryTree) Insert(i int64) BinaryTree {
	cbt.Lock()
	defer cbt.Unlock()
	if cbt.root == nil {
		cbt.root = initTreeNode(i)
		cbt.curLen++
		return cbt
	}
	cbt.root.insert(i)
	cbt.curLen++
	return cbt
}

func (cbt *concurrencyBinaryTree) Inserts(l []int64) BinaryTree {
	cbt.Lock()
	defer cbt.Unlock()
	if len(l) == 0 {
		return cbt
	}
	if cbt.root == nil {
		cbt.root = initTreeNode(l[0])
	}
	for i := 1; i < len(l); i++ {
		cbt.root.insert(l[i])
	}
	return cbt
}

func (cbt *concurrencyBinaryTree) Search(i int64) int {
	cbt.RLock()
	defer cbt.RUnlock()
	if cbt.root == nil {
		return 0
	}
	return cbt.root.search(i)
}

func (cbt *concurrencyBinaryTree) Depth() int {
	cbt.RLock()
	defer cbt.RUnlock()
	if cbt.root == nil {
		return 0
	}
	return cbt.root.depth()
}

func (cbt *concurrencyBinaryTree) Len() int {
	cbt.RLock()
	defer cbt.RUnlock()
	return cbt.len()
}

func (cbt *concurrencyBinaryTree) len() int {
	return cbt.curLen
}

func (cbt *concurrencyBinaryTree) InorderTraversal() (res []int64) {
	cbt.RLock()
	defer cbt.RUnlock()
	if cbt.root == nil {
		return nil
	}
	cbt.root.inorder(&res)
	return res
}

func (cbt *concurrencyBinaryTree) PreorderTraversal() (res []int64) {
	cbt.RLock()
	defer cbt.RUnlock()
	if cbt.root == nil {
		return nil
	}
	cbt.root.preorder(&res)
	return res
}

func (cbt *concurrencyBinaryTree) PostorderTraversal() (res []int64) {
	cbt.RLock()
	defer cbt.RUnlock()
	if cbt.root == nil {
		return nil
	}
	cbt.root.postorder(&res)
	return res
}

func (cbt *concurrencyBinaryTree) Delete(i int64) int {
	cbt.Lock()
	defer cbt.Unlock()
	if cbt.root == nil {
		return 0
	}
	parent, cur := cbt.root, cbt.root
	isLeft := true
	for {
		if i < cur.val {
			if cur.left != nil {
				if cur.left.val == i {
					parent = cur
					cur = cur.left
					break
				}
				cur = cur.left
				continue
			}
		} else if i > cur.val {
			if cur.right != nil {
				if cur.right.val == i {
					parent = cur
					cur = cur.right
					isLeft = false
					break
				}
				cur = cur.right
				continue
			}
		} else {
			break
		}
		return 0
	}
	// 要删除的节点是叶子节点
	if cur.left == nil && cur.right == nil {
		if cur == cbt.root {
			// 特殊情况，只有根节点且根节点就被删除
			cbt.root = nil
		} else if isLeft {
			parent.left = nil
		} else {
			parent.right = nil
		}
		// 要删的节点只有左节点
	} else if cur.right == nil {
		if cur == cbt.root {
			cbt.root = cur.left
		} else if isLeft {
			parent.left = cur.left
		} else {
			parent.right = cur.left
		}
		// 要删的节点只有右节点
	} else if cur.left == nil {
		if cur == cbt.root {
			cbt.root = cur.right
		} else if isLeft {
			parent.left = cur.right
		} else {
			parent.right = cur.right
		}
	} else {
		// 最复杂的情况，既有左节点又有右节点
		suc := cur.getSuccessor()
		if cur == cbt.root {
			cbt.root = suc
		} else if isLeft {
			parent.left = suc
		} else {
			parent.right = suc
		}
		suc.left = cur.left
	}
	return cur.count
}
