func searchBST(root *TreeNode, val int) *TreeNode {
	for root != nil {
		switch {
		case root.Val > val:
			root = root.Left
		case root.Val < val:
			root = root.Right
		default:
			return root
		}
	}
	return nil
}