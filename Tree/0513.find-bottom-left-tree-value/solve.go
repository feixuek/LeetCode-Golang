func findBottomLeftValue(root *TreeNode) int {
	var res int
	quene := []*TreeNode{root}
	end := len(quene)
	for end != 0 {
		for i := 0; i < end; i++ {
			elem := quene[i]
			if elem.Left != nil {
				quene = append(quene, elem.Left)
			}
			if elem.Right != nil {
				quene = append(quene, elem.Right)
			}
		}
		res = quene[0].Val
		quene = quene[end:]
		end = len(quene)
	}
	return res
}