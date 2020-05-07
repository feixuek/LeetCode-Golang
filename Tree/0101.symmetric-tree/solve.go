func isSymmetric(root *TreeNode) bool {
	return isMirror(root, root)
}

func isMirror(t, n *TreeNode) bool {
	if t == nil && n == nil {
		return true
	}
	if (t == nil && n != nil) || (t != nil && n == nil) {
		return false
	}
	return (t.Val == n.Val) && isMirror(t.Left, n.Right) && isMirror(t.Right, n.Left)
}