func deepestLeavesSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	quene := []*TreeNode{root}
	end := len(quene)
	res := 0
	for end != 0 {
		sum := 0
		for i := 0; i < end; i++ {
			elem := quene[i]
			sum += elem.Val
			if elem.Left != nil {
				quene = append(quene, elem.Left)
			}
			if elem.Right != nil {
				quene = append(quene, elem.Right)
			}
		}
		quene = quene[end:]
		end = len(quene)
		res = sum
	}
	return res
}