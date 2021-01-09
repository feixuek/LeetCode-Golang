func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	switch {
	case left != nil && right != nil:
		return root
	case left != nil:
		return left
	case right != nil:
		return right
	default:
		return nil
	}
}