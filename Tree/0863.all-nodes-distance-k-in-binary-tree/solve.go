func distanceK(root *TreeNode, target *TreeNode, K int) []int {
	if root == nil {
		return nil
	}
	filter := make(map[int]int)
	_ = find(root, target, &filter)
	dis, ok := filter[root.Val]
	if !ok {
		return nil
	}
	var res []int
	dfs(root, K, dis, &res, &filter)
	return res
}

func find(root, target *TreeNode, filter *map[int]int) int {
	if root == nil {
		return -1
	}
	if root.Val == target.Val {
		(*filter)[root.Val] = 0
		return 0
	}
	left := find(root.Left, target, filter)
	if left >= 0 {
		(*filter)[root.Val] = left + 1
		return left + 1
	}
	right := find(root.Right, target, filter)
	if right >= 0 {
		(*filter)[root.Val] = right + 1
		return right + 1
	}
	return -1
}

func dfs(root *TreeNode, K, dis int, res *[]int, filter *map[int]int) {
	if root == nil {
		return
	}
	targetDis, ok := (*filter)[root.Val]
	if ok {
		dis = targetDis
	}
	if dis == K {
		(*res) = append(*res, root.Val)
	}
	dfs(root.Left, K, dis+1, res, filter)
	dfs(root.Right, K, dis+1, res, filter)
}