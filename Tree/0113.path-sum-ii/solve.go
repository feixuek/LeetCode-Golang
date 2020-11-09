func pathSum(root *TreeNode, sum int) [][]int {
	cache := []int{}
	ret := [][]int{}
	helper(root, sum, &cache, &ret)
	return ret
}

func helper(root *TreeNode, sum int, cache *[]int, ret *[][]int) {
	if root == nil {
		return
	}
	*cache = append(*cache, root.Val)
	sum -= root.Val
	if root.Left == nil && root.Right == nil && sum == 0 {
		tmp := make([]int, len(*cache))
		copy(tmp, *cache)
		*ret = append(*ret, tmp)
	}
	helper(root.Left, sum, cache, ret)
	helper(root.Right, sum, cache, ret)
	(*cache) = (*cache)[:len(*cache)-1]
}