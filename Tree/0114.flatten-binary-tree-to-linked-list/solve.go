func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	left := root.Left
	root.Left = nil
	tmp := left
	if tmp != nil {
		for tmp.Right != nil {
			tmp = tmp.Right
		}
		tmp.Right = root.Right
		root.Right = left
	}
	flatten(root.Right)
}