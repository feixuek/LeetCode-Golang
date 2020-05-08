func preorder(root *Node) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0)
	res = append(res, root.Val)
	for i := 0; i < len(root.Children); i++ {
		res = append(res, preorder(root.Children[i])...)
	}
	return res
}