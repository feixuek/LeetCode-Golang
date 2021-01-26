func goodNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := new(int)
	helper(root, root.Val, res)
	return *res
}

func helper(root *TreeNode, max int, res *int) {
	if root == nil {
		return
	}
	if root.Val >= max {
		(*res)++
		max = root.Val
	}
	helper(root.Left, max, res)
	helper(root.Right, max, res)
}