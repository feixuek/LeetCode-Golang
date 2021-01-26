func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var pre *TreeNode
	stack := []*TreeNode{root}
	res := []int{}
	for len(stack) != 0 {
		curr := stack[len(stack)-1]
		if (curr.Left == nil && curr.Right == nil) || (pre != nil && (curr.Left == pre || curr.Right == pre)) {
			res = append(res, curr.Val)
			stack = stack[:len(stack)-1]
			pre = curr
		} else {
			if curr.Right != nil {
				stack = append(stack, curr.Right)
			}
			if curr.Left != nil {
				stack = append(stack, curr.Left)
			}
		}
	}
	return res
}