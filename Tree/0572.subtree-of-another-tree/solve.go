func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil {
		return false
	}
	if t == nil {
		return true
	}
	if isSame(s, t) {
		return true
	}
	return isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func isSame(s *TreeNode, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}
	if s == nil && t != nil {
		return false
	}
	if s != nil && t == nil {
		return false
	}
	if s.Val != t.Val {
		return false
	}
	return isSame(s.Left, t.Left) && isSame(s.Right, t.Right)
}