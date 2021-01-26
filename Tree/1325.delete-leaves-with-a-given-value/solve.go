func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	if root == nil {
		return nil
	}
	tmp := &TreeNode{
		Val:   -1,
		Left:  root,
		Right: nil,
	}
	helper(tmp.Left, tmp, target)
	return tmp.Left
}

func helper(root, pre *TreeNode, target int) {
	if root == nil {
		return
	}
	helper(root.Left, root, target)
	helper(root.Right, root, target)
	if root.Left == nil && root.Right == nil && root.Val == target {
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