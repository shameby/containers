package containers

type binaryTree struct {
	root   *binaryTreeNode
	curLen int
}

func initTreeNode(i int64) *binaryTreeNode {
	return &binaryTreeNode{
		i, 1, nil, nil,
	}
}

func (bt *binaryTree) Insert(i int64) BinaryTree {
	if bt.root == nil {
		bt.root = initTreeNode(i)
		bt.curLen++
		return bt
	}
	bt.root.insert(i)
	bt.curLen++
	return bt
}

func (bt *binaryTree) Inserts(l []int64) BinaryTree {
	if len(l) == 0 {
		return bt
	}
	if bt.root == nil {
		bt.root = initTreeNode(l[0])
	}
	for i := 1; i < len(l); i++ {
		bt.root.insert(l[i])
	}
	return bt
}

func (bt *binaryTree) Search(i int64) int {
	if bt.root == nil {
		return 0
	}
	return bt.root.search(i)
}

func (bt *binaryTree) Depth() int {
	if bt.root == nil {
		return 0
	}
	return bt.root.depth()
}

func (bt *binaryTree) Len() int {
	return bt.curLen
}

func (bt *binaryTree) InorderTraversal() (res []int64) {
	if bt.root == nil {
		return nil
	}
	bt.root.inorder(&res)
	return res
}

func (bt *binaryTree) PreorderTraversal() (res []int64) {
	if bt.root == nil {
		return nil
	}
	bt.root.preorder(&res)
	return res
}

func (bt *binaryTree) PostorderTraversal() (res []int64) {
	if bt.root == nil {
		return nil
	}
	bt.root.postorder(&res)
	return res
}

func (bt *binaryTree) Delete(i int64) int {
	if bt.root == nil {
		return 0
	}
	parent, cur := bt.root, bt.root
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
		if cur == bt.root {
			// 特殊情况，只有根节点且根节点就被删除
			bt.root = nil
		} else if isLeft {
			parent.left = nil
		} else {
			parent.right = nil
		}
		// 要删的节点只有左节点
	} else if cur.right == nil {
		if cur == bt.root {
			bt.root = cur.left
		} else if isLeft {
			parent.left = cur.left
		} else {
			parent.right = cur.left
		}
		// 要删的节点只有右节点
	} else if cur.left == nil {
		if cur == bt.root {
			bt.root = cur.right
		} else if isLeft {
			parent.left = cur.right
		} else {
			parent.right = cur.right
		}
	} else {
		// 最复杂的情况，既有左节点又有右节点
		suc := cur.getSuccessor()
		if cur == bt.root {
			bt.root = suc
		} else if isLeft {
			parent.left = suc
		} else {
			parent.right = suc
		}
		suc.left = cur.left
	}
	return cur.count
}
