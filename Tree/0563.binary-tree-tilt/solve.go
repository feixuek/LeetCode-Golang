func findTilt(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return treeTilt(root) + findTilt(root.Left) + findTilt(root.Right)
}

func treeTilt(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := treeSum(root.Left)
	right := treeSum(root.Right)
	if left > right {
		return left - right
	}
	return right - left
}

func treeSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return root.Val + treeSum(root.Left) + treeSum(root.Right)
}