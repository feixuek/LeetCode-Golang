func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	quene := []*TreeNode{root}
	end := len(quene)
	res := make([][]int, 0)
	for end != 0 {
		data := []int{}
		for i := 0; i < end; i++ {
			elem := quene[i]
			data = append(data, elem.Val)
			if elem.Left != nil {
				quene = append(quene, elem.Left)
			}
			if elem.Right != nil {
				quene = append(quene, elem.Right)
			}
		}
		quene = quene[end:]
		end = len(quene)
		res = append(res, data)
	}
	left := 0
	right := len(res) - 1
	for left < right {
		res[left], res[right] = res[right], res[left]
		left++
		right--
	}
	return res
}