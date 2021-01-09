func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	if root == nil {
		return nil
	}
	filter := make(map[int]struct{})
	for i := 0; i < len(to_delete); i++ {
		filter[to_delete[i]] = struct{}{}
	}
	res := make([]*TreeNode, 0)
	find(root, filter, &res, true)
	return res
}

func find(root *TreeNode, filter map[int]struct{}, res *[]*TreeNode, isRoot bool) *TreeNode {
	if root == nil {
		return nil
	}
	needDelete := false
	if _, ok := filter[root.Val]; ok {
		needDelete = true
	}
	if !needDelete && isRoot {
		(*res) = append(*res, root)
	}
	root.Left = find(root.Left, filter, res, needDelete)
	root.Right = find(root.Right, filter, res, needDelete)
	if needDelete {
		return nil
	}
	return root
}