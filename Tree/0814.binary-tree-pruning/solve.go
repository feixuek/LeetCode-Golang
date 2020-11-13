func pruneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if needPrune(root) {
		return nil
	}
	root.Left = pruneTree(root.Left)
	root.Right = pruneTree(root.Right)
	return root
}

func needPrune(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return root.Val == 0 && needPrune(root.Left) && needPrune(root.Right)
}