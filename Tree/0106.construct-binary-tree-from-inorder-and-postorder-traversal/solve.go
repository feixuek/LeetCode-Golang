func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 || len(inorder) != len(postorder) {
		return nil
	}
	return helper(inorder, postorder, 0, len(inorder)-1, 0, len(postorder)-1)
}

func helper(inorder []int, postorder []int, startIn, endIn, startPost, endPost int) *TreeNode {
	if startIn > endIn || startPost > endPost {
		return nil
	}
	elem := postorder[endPost]
	idx := startIn
	for i := startIn; i <= endIn; i++ {
		if inorder[i] == elem {
			idx = i
			break
		}
	}
	ret := &TreeNode{
		Val:   elem,
		Left:  helper(inorder, postorder, startIn, idx-1, startPost, startPost+idx-startIn-1),
		Right: helper(inorder, postorder, idx+1, endIn, startPost+idx-startIn, endPost-1),
	}
	return ret
}