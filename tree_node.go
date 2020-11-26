package containers

type binaryTreeNode struct {
	val   *Elem
	left  *binaryTreeNode
	right *binaryTreeNode
}

func initTreeNode(elem *Elem) *binaryTreeNode {
	return &binaryTreeNode{
		elem, nil, nil,
	}
}

func (ntn *binaryTreeNode) insert(elem *Elem) {
	switch {
	case elem.Score < ntn.val.Score:
		if ntn.left == nil {
			ntn.left = initTreeNode(elem)
			return
		}
		ntn.left.insert(elem)
	case elem.Score >= ntn.val.Score:
		if ntn.right == nil {
			ntn.right = initTreeNode(elem)
			return
		}
		ntn.right.insert(elem)
	}
}

func (ntn *binaryTreeNode) search(i float64) *Elem {
	switch {
	case i < ntn.val.Score:
		if ntn.left != nil {
			return ntn.left.search(i)
		}
	case i > ntn.val.Score:
		if ntn.right != nil {
			return ntn.right.search(i)
		}
	default:
		return ntn.val
	}
	return nil
}

func (ntn *binaryTreeNode) depth() int {
	if ntn != nil {
		return maxInt(ntn.left.depth(), ntn.right.depth()) + 1
	}
	return 0
}

func (ntn *binaryTreeNode) inorder(res *[]Elem) {
	if ntn != nil {
		ntn.left.inorder(res)
		*res = append(*res, *ntn.val)
		ntn.right.inorder(res)
	}
}

func (ntn *binaryTreeNode) preorder(res *[]Elem) {
	if ntn != nil {
		*res = append(*res, *ntn.val)
		ntn.left.preorder(res)
		ntn.right.preorder(res)
	}
}

func (ntn *binaryTreeNode) postorder(res *[]Elem) {
	if ntn != nil {
		ntn.left.postorder(res)
		ntn.right.postorder(res)
		*res = append(*res, *ntn.val)
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