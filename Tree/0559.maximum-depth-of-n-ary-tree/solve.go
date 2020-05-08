func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	if len(root.Children) == 0 {
		return 1
	}
	depth := make([]int, len(root.Children))
	for i := 0; i < len(root.Children); i++ {
		depth[i] = maxDepth(root.Children[i])
	}
	max := depth[0]
	for i := 1; i < len(depth); i++ {
		if depth[i] > max {
			max = depth[i]
		}
	}
	return max + 1
}