func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	left := depth(root.Left)
	right := depth(root.Right)
	c := left - right
	if c > 1 || c < (-1) {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := depth(root.Left)
	right := depth(root.Right)
	if left >= right {
		return left + 1
	}
	return right + 1
}