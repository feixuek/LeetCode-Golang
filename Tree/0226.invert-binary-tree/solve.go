func invertTree(root *TreeNode) *TreeNode {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return root
	}
	ret := &TreeNode{
		Val: root.Val,
	}
	ret.Left = invertTree(root.Right)
	ret.Right = invertTree(root.Left)
	return ret
}