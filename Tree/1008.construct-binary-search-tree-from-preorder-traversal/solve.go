func bstFromPreorder(preorder []int) *TreeNode {
	return helper(preorder, 0, len(preorder)-1)
}

func helper(preorder []int, start, end int) *TreeNode {
	if start > end {
		return nil
	}
	root := &TreeNode{
		Val:   preorder[start],
		Left:  nil,
		Right: nil,
	}
	index := start
	for i := start + 1; i <= end; i++ {
		if preorder[i] > preorder[start] {
			break
		}
		index++
	}
	root.Left = helper(preorder, start+1, index)
	root.Right = helper(preorder, index+1, end)
	return root
}