func convertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	sum := new(int)
	convert(root, sum)
	return root
}

func convert(root *TreeNode, sum *int) {
	if root == nil {
		return
	}
	convert(root.Right, sum)
	root.Val += *sum
	*sum = root.Val
	convert(root.Left, sum)
}
