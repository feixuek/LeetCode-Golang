func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	left := minDepth(root.Left)
	right := minDepth(root.Right)
	if root.Left == nil || root.Right == nil {
		return left + right + 1
	}
	if left < right {
		return left + 1
	}
	return right + 1
}