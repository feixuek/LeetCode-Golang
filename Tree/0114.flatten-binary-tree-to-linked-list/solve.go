func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Left)
	flatten(root.Right)
	left := root.Left
	right := root.Right
	root.Left = nil
	root.Right = left
	if root.Right != nil {
		target := root.Right
		for target != nil && target.Right != nil {
			target = target.Right
		}
		target.Right = right
	} else {
		root.Right = right
	}
}