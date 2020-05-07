func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		return root.Left.Val + sumOfLeftLeaves(root.Right)
	}
	left := sumOfLeftLeaves(root.Left)
	right := sumOfLeftLeaves(root.Right)
	return left + right
}