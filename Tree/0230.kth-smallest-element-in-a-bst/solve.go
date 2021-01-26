func kthSmallest(root *TreeNode, k int) int {
	res := new(int)
	count := new(int)
	*count = k
	helper(root, count, res)
	return *res
}

func helper(root *TreeNode, k, res *int) {
	if root == nil {
		return
	}
	helper(root.Left, k, res)
	(*k)--
	if (*k) == 0 {
		*res = root.Val
	}
	helper(root.Right, k, res)
}