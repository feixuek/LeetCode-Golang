func sumNumbers(root *TreeNode) int {
	path := new(int)
	sum := new(int)
	helper(root, path, sum)
	return *sum
}

func helper(root *TreeNode, path, sum *int) {
	if root == nil {
		return
	}
	*path = (*path)*10 + root.Val
	if root.Left == nil && root.Right == nil {
		*sum += *path
	}
	helper(root.Left, path, sum)
	helper(root.Right, path, sum)
	*path = ((*path) - root.Val) / 10
}