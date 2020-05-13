func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return 0
	}
	diameter := maxDepth(root.Left) + maxDepth(root.Right)
	return max(diameter, max(diameterOfBinaryTree(root.Left), diameterOfBinaryTree(root.Right)))
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	return max(left, right) + 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}