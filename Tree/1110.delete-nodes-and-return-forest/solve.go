func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	if root == nil {
		return nil
	}
	filter := make(map[int]bool)
	for i := 0; i < len(to_delete); i++ {
		filter[to_delete[i]] = true
	}
	res := []*TreeNode{}
	helper(root, nil, filter, true, &res)
	return res
}

func helper(root, pre *TreeNode, filter map[int]bool, isRoot bool, res *[]*TreeNode) {
	if root == nil {
		return
	}
	ok := filter[root.Val]
	if ok {
		if pre != nil {
			if pre.Left == root {
				pre.Left = nil
			}
			if pre.Right == root {
				pre.Right = nil
			}
		}
	} else {
		if isRoot {
			*res = append(*res, root)
		}
	}
	helper(root.Left, root, filter, ok, res)
	helper(root.Right, root, filter, ok, res)
}