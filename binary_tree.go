package containers

type binaryTree struct {
	root   *binaryTreeNode
	curLen int
	preL   []Elem
	inL    []Elem
	postL  []Elem
}

func (bt *binaryTree) Insert(elem IElem) BinaryTree {
	bt.initL()
	e := initE(elem)
	if bt.root == nil {
		bt.root = initTreeNode(e)
		bt.curLen++
		return bt
	}
	bt.root.insert(e)
	bt.curLen++
	return bt
}

func (bt *binaryTree) Search(i int64) *Elem {
	if bt.root == nil {
		return nil
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

func (bt *binaryTree) InorderTraversal() (res []Elem) {
	if bt.root == nil {
		return nil
	}
	if bt.inL == nil {
		bt.inL = make([]Elem, 0)
		bt.root.inorder(&bt.inL)
	}
	return bt.inL
}

func (bt *binaryTree) PreorderTraversal() (res []Elem) {
	if bt.root == nil {
		return nil
	}
	if bt.preL == nil {
		bt.preL = make([]Elem, 0)
		bt.root.preorder(&bt.preL)
	}
	return bt.preL
}

func (bt *binaryTree) PostorderTraversal() []Elem {
	if bt.root == nil {
		return nil
	}
	if bt.postL == nil {
		bt.postL = make([]Elem, 0)
		bt.root.postorder(&bt.postL)
	}
	return bt.postL
}

func (bt *binaryTree) Delete(i int64) int {
	if bt.root == nil {
		return 0
	}
	bt.initL()
	parent, cur := bt.root, bt.root
	isLeft := true
	for {
		if i < cur.val.Score {
			if cur.left != nil {
				if cur.left.val.Score == i {
					parent = cur
					cur = cur.left
					break
				}
				cur = cur.left
				continue
			}
		} else if i > cur.val.Score {
			if cur.right != nil {
				if cur.right.val.Score == i {
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
	return 1
}

func (bt *binaryTree) initL() {
	bt.preL, bt.inL, bt.postL = nil, nil, nil
}
