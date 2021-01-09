func constructFromPrePost(pre []int, post []int) *TreeNode {
	return helper(pre, post, 0, len(pre)-1, 0, len(post)-1)
}

func helper(pre, post []int, preStart, preEnd, postStart, postEnd int) *TreeNode {
	if preStart >= len(pre) || preStart > preEnd {
		return nil
	}
	root := &TreeNode{Val: pre[preStart]}
	if preStart == preEnd {
		return root
	}
	val := pre[preStart+1]
	step := 0
	for i := postStart; i < postEnd; i++ {
		if post[i] == val {
			break
		}
		step++
	}
	root.Left = helper(pre, post, preStart+1, preStart+1+step, postStart, postStart+step)
	root.Right = helper(pre, post, preStart+2+step, preEnd, postStart+step+1, postEnd-1)
	return root
}