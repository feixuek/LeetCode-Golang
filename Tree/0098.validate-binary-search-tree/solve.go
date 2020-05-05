func isValidBST(root *TreeNode) bool {
	max := 1 << 32
	min := -(max - 1)
	return isOk(root, min, max)
}

func isOk(node *TreeNode, min, max int) bool {
	if node == nil {
		return true
	}
	if node.Val <= min || node.Val >= max {
		return false
	}
	return isOk(node.Left, min, node.Val) && isOk(node.Right, node.Val, max)
}