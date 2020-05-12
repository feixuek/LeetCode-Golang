func isUnivalTree(root *TreeNode) bool {
	return helper(root, root.Val)
}

func helper(root *TreeNode, val int) bool {
	if root == nil {
		return true
	}
	if root.Val != val {
		return false
	}
	return helper(root.Left, val) && helper(root.Right, val)
}