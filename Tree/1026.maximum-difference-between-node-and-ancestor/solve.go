func maxAncestorDiff(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left, right := 0, 0
	if root.Left != nil {
		left = maxDiff(root.Val, root.Left)
	}
	if root.Right != nil {
		right = maxDiff(root.Val, root.Right)
	}
	return max(max(maxAncestorDiff(root.Left), maxAncestorDiff(root.Right)), max(left, right))
}

func maxDiff(val int, root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(val-root.Val, max(maxDiff(val, root.Left), maxDiff(val, root.Right)))
}

func max(x, y int) int {
	if x < 0 {
		x *= -1
	}
	if y < 0 {
		y *= -1
	}
	if x < y {
		return y
	}
	return x
}