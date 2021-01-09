	if root != nil && root.Val > val {
		root.Right = insertIntoMaxTree(root.Right, val)
		return root
	}
	return &TreeNode{
		Val:  val,
		Left: root,
	}