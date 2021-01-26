func maxAncestorDiff(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := new(int)
	helper(root.Left, root.Val, res)
	helper(root.Right, root.Val, res)
	return max(*res, max(maxAncestorDiff(root.Left), maxAncestorDiff(root.Right)))
}

func helper(root *TreeNode, val int, res *int) {
	if root == nil {
		return
	}
	target := root.Val - val
	if target < 0 {
		target *= -1
	}
	if target > *res {
		*res = target
	}
	helper(root.Left, val, res)
	helper(root.Right, val, res)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}