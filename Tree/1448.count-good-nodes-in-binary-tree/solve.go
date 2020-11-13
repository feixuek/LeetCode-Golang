func goodNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + help(root.Left, root.Val) + help(root.Right, root.Val)
}

func help(root *TreeNode, min int) int {
	if root == nil {
		return 0
	}
	ret := 0
	if root.Val >= min {
		min = root.Val
		ret = 1
	}
	return ret + help(root.Left, min) + help(root.Right, min)
}