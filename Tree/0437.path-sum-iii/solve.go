func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	return helper(root, sum) + pathSum(root.Left, sum) + pathSum(root.Right, sum)
}

func helper(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	ret := 0
	if root.Val == sum {
		ret = 1
	}
	return ret + helper(root.Left, sum-root.Val) + helper(root.Right, sum-root.Val)
}