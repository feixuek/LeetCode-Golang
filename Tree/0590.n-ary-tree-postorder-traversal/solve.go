func postorder(root *Node) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0)
	for i := 0; i < len(root.Children); i++ {
		res = append(res, postorder(root.Children[i])...)
	}
	res = append(res, root.Val)
	return res
}