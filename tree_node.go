package data_structures

type binaryTreeNode struct {
	val   int64
	count int
	left  *binaryTreeNode
	right *binaryTreeNode
}

func (ntn *binaryTreeNode) insert(i int64) {
	switch {
	case i < ntn.val:
		if ntn.left == nil {
			ntn.left = initTreeNode(i)
			return
		}
		ntn.left.insert(i)
	case i > ntn.val:
		if ntn.right == nil {
			ntn.right = initTreeNode(i)
			return
		}
		ntn.right.insert(i)
	default:
		ntn.count++
	}
}

func (ntn *binaryTreeNode) search(i int64) int {
	switch {
	case i < ntn.val:
		if ntn.left != nil {
			return ntn.left.search(i)
		}
	case i > ntn.val:
		if ntn.right != nil {
			return ntn.right.search(i)
		}
	default:
		return ntn.count
	}
	return 0
}

func (ntn *binaryTreeNode) inorder(res *[]int64) {
	if ntn != nil {
		ntn.left.inorder(res)
		for i := 1; i <= ntn.count; i++ {
			*res = append(*res, ntn.val)
		}
		ntn.right.inorder(res)
	}
}

func (ntn *binaryTreeNode) preorder(res *[]int64) {
	if ntn != nil {
		for i := 1; i <= ntn.count; i++ {
			*res = append(*res, ntn.val)
		}
		ntn.left.preorder(res)
		ntn.right.preorder(res)
	}
}

func (ntn *binaryTreeNode) postorder(res *[]int64) {
	if ntn != nil {
		ntn.left.postorder(res)
		ntn.right.postorder(res)
		for i := 1; i <= ntn.count; i++ {
			*res = append(*res, ntn.val)
		}
	}
}

func (ntn *binaryTreeNode) getSuccessor() *binaryTreeNode {
	var sucParent *binaryTreeNode = nil
	suc, cur := ntn.right, ntn.right
	for cur != nil {
		sucParent = suc
		suc = cur
		cur = cur.left
	}
	if suc != ntn.right {
		sucParent.left = suc.right
		suc.right = ntn.right
	}
	return suc
}