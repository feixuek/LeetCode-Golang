func findFrequentTreeSum(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	cache := make(map[int]int)
	_ = helper(root, &cache)
	var ret []int
	var number int
	for v, num := range cache {
		switch {
		case num > number || number == 0:
			ret = ret[:0]
			ret = append(ret, v)
			number = num
		case num == number:
			ret = append(ret, v)
		}
	}
	return ret
}

func helper(root *TreeNode, cache *map[int]int) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		(*cache)[root.Val]++
		return root.Val
	}
	left := helper(root.Left, cache)
	right := helper(root.Right, cache)
	sum := right + left + root.Val
	(*cache)[sum]++
	return sum
}