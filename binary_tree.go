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
