func pruneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	pre := &TreeNode{
		Left: root,
	}
	helper(root, pre)
	return pre.Left
}

func helper(root, pre *TreeNode) {
	if root == nil {
		return
	}
	helper(root.Left, root)
	helper(root.Right, root)
	if root.Left == nil && root.Right == nil && root.Val == 0 {
		if pre != nil {
			if pre.Left == root {
				pre.Left = nil
			}
			if pre.Right == root {
				pre.Right = nil
			}
		}
	}
}