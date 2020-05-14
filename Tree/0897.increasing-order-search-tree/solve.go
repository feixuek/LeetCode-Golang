func increasingBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	return increasingBSTHelper(root, nil)
}

func increasingBSTHelper(root, tail *TreeNode) *TreeNode {
	if root == nil {
		return tail
	}
	res := increasingBSTHelper(root.Left, root)
	root.Left = nil
	root.Right = increasingBSTHelper(root.Right, tail)
	return res
}