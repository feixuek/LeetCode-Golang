func longestUnivaluePath(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ret := new(int)
	maxdepth(root, root.Val, ret)
	return *ret
}

func maxdepth(root *TreeNode, val int, ret *int) int {
	if root == nil {
		return 0
	}
	left := maxdepth(root.Left, root.Val, ret)
	right := maxdepth(root.Right, root.Val, ret)
	sum := left + right
	if sum > *ret {
		*ret = sum
	}
	if root.Val == val {
		if left > right {
			return left + 1
		}
		return right + 1
	}
	return 0
}