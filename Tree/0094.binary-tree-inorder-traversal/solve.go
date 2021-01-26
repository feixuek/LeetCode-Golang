func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	stack := []*TreeNode{}
	curr := root
	for curr != nil || len(stack) != 0 {
		if curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		} else {
			elem := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = append(res, elem.Val)
			curr = elem.Right
		}
	}
	return res
}