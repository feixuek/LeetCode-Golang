func largestValues(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	quene := []*TreeNode{root}
	end := len(quene)
	res := []int{}
	for end != 0 {
		max := -(1 << 31)
		for i := 0; i < end; i++ {
			elem := quene[i]
			if elem.Val > max {
				max = elem.Val
			}
			if elem.Left != nil {
				quene = append(quene, elem.Left)
			}
			if elem.Right != nil {
				quene = append(quene, elem.Right)
			}
		}
		res = append(res, max)
		quene = quene[end:]
		end = len(quene)
	}
	return res
}