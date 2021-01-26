func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	quene := []*TreeNode{root}
	res := []int{}
	end := len(quene)
	for len(quene) != 0 {
		for i := 0; i < end; i++ {
			elem := quene[i]
			if elem.Left != nil {
				quene = append(quene, elem.Left)
			}
			if elem.Right != nil {
				quene = append(quene, elem.Right)
			}
		}
		res = append(res, quene[end-1].Val)
		quene = quene[end:]
		end = len(quene)
	}
	return res
}