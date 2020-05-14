func rangeSumBST(root *TreeNode, L int, R int) int {
	if root == nil {
		return 0
	}
	switch {
	case root.Val >= L && root.Val <= R:
		return root.Val + rangeSumBST(root.Left, L, R) + rangeSumBST(root.Right, L, R)
	case root.Val < L:
		return rangeSumBST(root.Right, L, R)
	case root.Val > R:
		return rangeSumBST(root.Left, L, R)
	}
	return 0
}