func findTarget(root *TreeNode, k int) bool {
	return helper(root, root, k)
}

func helper(root, elem *TreeNode, k int) bool {
	if elem == nil {
		return false
	}
	return findElem(root, k-elem.Val, elem.Val) || helper(root, elem.Left, k) || helper(root, elem.Right, k)
}

func findElem(root *TreeNode, elem, filter int) bool {
	for root != nil {
		switch {
		case root.Val == elem:
			if root.Val != filter {
				return true
			}
			return false
		case root.Val > elem:
			root = root.Left
		default:
			root = root.Right
		}
	}
	return false
}