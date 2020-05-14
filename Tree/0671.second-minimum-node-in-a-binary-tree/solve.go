func findSecondMinimumValue(root *TreeNode) int {
	if root == nil {
		return -1
	}
	return helper(root, root.Val)
}

func helper(root *TreeNode, val int) int {
	if root == nil {
		return -1
	}
	if root.Val != val {
		return root.Val
	}
	left := helper(root.Left, val)
	right := helper(root.Right, val)
	if left == -1 {
		return right
	}
	if right == -1 {
		return left
	}
	if left < right {
		return left
	}
	return right
}