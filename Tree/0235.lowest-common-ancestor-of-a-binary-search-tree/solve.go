func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	for ((root.Val - p.Val) * (root.Val - q.Val)) > 0 {
		if p.Val < root.Val {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return root
}