func sumEvenGrandparent(root *TreeNode) int {
	sum := new(int)
	help(root, sum)
	return *sum
}

func help(root *TreeNode, sum *int) {
	if root == nil {
		return
	}
	if root.Val%2 == 0 {
		if root.Left != nil {
			if root.Left.Left != nil {
				*sum = *sum + root.Left.Left.Val
			}
			if root.Left.Right != nil {
				*sum = *sum + root.Left.Right.Val
			}
		}
		if root.Right != nil {
			if root.Right.Left != nil {
				*sum = *sum + root.Right.Left.Val
			}
			if root.Right.Right != nil {
				*sum = *sum + root.Right.Right.Val
			}
		}
	}
	help(root.Left, sum)
	help(root.Right, sum)
}