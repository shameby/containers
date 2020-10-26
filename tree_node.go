package data_structures

type binaryTreeNode struct {
	val   int64
	count int
	left  *binaryTreeNode
	right *binaryTreeNode
}

func (tn *binaryTreeNode) insert(i int64) {
	switch {
	case i < tn.val:
		if tn.left == nil {
			tn.left = initTreeNode(i)
			return
		}
		tn.left.insert(i)
	case i > tn.val:
		if tn.right == nil {
			tn.right = initTreeNode(i)
			return
		}
		tn.right.insert(i)
	default:
		tn.count++
	}
}

func (tn *binaryTreeNode) inorder(res *[]int64) {
	if tn == nil {
		return
	}
	tn.left.inorder(res)
	for i := 1; i <= tn.count; i++ {
		*res = append(*res, tn.val)
	}
	tn.right.inorder(res)
}

func (tn *binaryTreeNode) preorder(res *[]int64) {
	if tn == nil {
		return
	}
	for i := 1; i <= tn.count; i++ {
		*res = append(*res, tn.val)
	}
	tn.left.preorder(res)
	tn.right.preorder(res)
}

func (tn *binaryTreeNode) postorder(res *[]int64) {
	if tn == nil {
		return
	}
	tn.left.postorder(res)
	tn.right.postorder(res)
	for i := 1; i <= tn.count; i++ {
		*res = append(*res, tn.val)
	}
}
