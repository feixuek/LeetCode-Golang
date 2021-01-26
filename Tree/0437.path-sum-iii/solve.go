func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	res := new(int)
	helper(root, sum, res)
	return *res + pathSum(root.Left, sum) + pathSum(root.Right, sum)
}

func helper(root *TreeNode, sum int, res *int) {
	if root == nil {
		return
	}
	sum -= root.Val
	if sum == 0 {
		(*res)++
	}
	helper(root.Left, sum, res)
	helper(root.Right, sum, res)
}