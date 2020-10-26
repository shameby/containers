package data_structures

type normalBinaryTree struct {
	root *binaryTreeNode
	len  int
}

func NewBinaryTree() BinaryTree {
	return &normalBinaryTree{
		nil, 0,
	}
}

func initTreeNode(i int64) *binaryTreeNode {
	return &binaryTreeNode{
		i, 1, nil, nil,
	}
}

func (nbt *normalBinaryTree) Insert(i int64) *normalBinaryTree {
	if nbt.root == nil {
		nbt.root = initTreeNode(i)
		nbt.len++
		return nbt
	}
	nbt.root.insert(i)
	nbt.len++
	return nbt
}

func (nbt *normalBinaryTree) Inserts(l []int64) *normalBinaryTree {
	if len(l) == 0 {
		return nbt
	}
	if nbt.root == nil {
		nbt.root = initTreeNode(l[0])
	}
	for i := 1; i < len(l); i++ {
		nbt.root.insert(l[i])
	}
	return nbt
}

func (nbt *normalBinaryTree) Search(i int64) int {
	if nbt.root == nil {
		return 0
	}
	return nbt.root.search(i)
}

func (nbt *normalBinaryTree) Len() int {
	return nbt.len
}

func (nbt *normalBinaryTree) InorderTraversal() (res []int64) {
	if nbt.root == nil {
		return nil
	}
	nbt.root.inorder(&res)
	return res
}

func (nbt *normalBinaryTree) PreorderTraversal() (res []int64) {
	if nbt.root == nil {
		return nil
	}
	nbt.root.preorder(&res)
	return res
}

func (nbt *normalBinaryTree) PostorderTraversal() (res []int64) {
	if nbt.root == nil {
		return nil
	}
	nbt.root.postorder(&res)
	return res
}

func (nbt *normalBinaryTree) Delete(i int64) int {
	if nbt.root == nil {
		return 0
	}
	parent, cur := nbt.root, nbt.root
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
		if cur == nbt.root {
			// 特殊情况，只有根节点且根节点就被删除
			nbt.root = nil
		} else if isLeft {
			parent.left = nil
		} else {
			parent.right = nil
		}
		// 要删的节点只有左节点
	} else if cur.right == nil {
		if cur == nbt.root {
			nbt.root = cur.left
		} else if isLeft {
			parent.left = cur.left
		} else {
			parent.right = cur.left
		}
		// 要删的节点只有右节点
	} else if cur.left == nil {
		if cur == nbt.root {
			nbt.root = cur.right
		} else if isLeft {
			parent.left = cur.right
		} else {
			parent.right = cur.right
		}
	} else {
		// 最复杂的情况，既有左节点又有右节点
		suc := cur.getSuccessor()
		if cur == nbt.root {
			nbt.root = suc
		} else if isLeft {
			parent.left = suc
		} else {
			parent.right = suc
		}
		suc.left = cur.left
	}
	return cur.count
}
