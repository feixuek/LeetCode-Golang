func longestZigZag(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ret, length := 0, 0
	help(root, false, length, &ret)
	help(root, true, length, &ret)
	return ret
}

func help(root *TreeNode, gotoLeft bool, length int, ret *int) {
	if length > *ret {
		*ret = length
	}
	if gotoLeft {
		if root.Left != nil {
			help(root.Left, !gotoLeft, length+1, ret)
		}
		if root.Right != nil {
			help(root.Right, gotoLeft, 1, ret)
		}
	} else {
		if root.Right != nil {
			help(root.Right, !gotoLeft, length+1, ret)
		}
		if root.Left != nil {
			help(root.Left, gotoLeft, 1, ret)
		}
	}
}