func findMode(root *TreeNode) []int {
	data := make(map[int]int)
	helper(root, data)
	maxCount := 0
	res := []int{}
	for val, num := range data {
		switch {
		case num == maxCount:
			if maxCount == 0 {
				maxCount = num
			}
			res = append(res, val)
		case num > maxCount:
			maxCount = num
			if len(res) > 0 {
				res = res[0:0]
			}
			res = append(res, val)
		}
	}
	return res
}

func helper(root *TreeNode, data map[int]int) {
	if root == nil {
		return
	}
	helper(root.Left, data)
	data[root.Val]++
	helper(root.Right, data)
}