func sumNumbers(root *TreeNode) int {
	return helper(root, 0)
}

func helper(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return root.Val + sum*10
	}
	sum = sum*10 + root.Val
	return helper(root.Left, sum) + helper(root.Right, sum)
}