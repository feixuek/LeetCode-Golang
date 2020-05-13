func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	tree := new(TreeNode)
	tree.Val = t1.Val + t2.Val
	tree.Left = mergeTrees(t1.Left, t2.Left)
	tree.Right = mergeTrees(t1.Right, t2.Right)
	return tree
}