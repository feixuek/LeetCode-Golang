func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 || len(inorder) != len(preorder) {
		return nil
	}
	return helper(preorder, inorder, 0, len(preorder)-1, 0, len(inorder)-1)
}

func helper(preorder []int, inorder []int, startPre, endPre, startIn, endIn int) *TreeNode {
	if startPre > endPre || startIn > endIn {
		return nil
	}
	elem := preorder[startPre]
	idx := startIn
	for i := startIn; i <= endIn; i++ {
		if inorder[i] == elem {
			idx = i
			break
		}
	}
	ret := &TreeNode{
		Val:   elem,
		Left:  helper(preorder, inorder, startPre+1, startPre+idx-startIn, startIn, idx-1),
		Right: helper(preorder, inorder, startPre+idx-startIn+1, endPre, idx+1, endIn),
	}
	return ret
}