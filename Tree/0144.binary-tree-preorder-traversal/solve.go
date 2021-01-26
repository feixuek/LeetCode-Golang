func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	stack := []*TreeNode{root}
	res := []int{}
	for len(stack) != 0 {
		elem := stack[len(stack)-1]
		res = append(res, elem.Val)
		stack = stack[:len(stack)-1]
		if elem.Right != nil {
			stack = append(stack, elem.Right)
		}
		if elem.Left != nil {
			stack = append(stack, elem.Left)
		}
	}
	return res
}