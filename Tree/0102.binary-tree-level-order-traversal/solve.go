func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	quene := []*TreeNode{root}
	res := make([][]int, 0)
	end := len(quene)
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
		res = append(res, data)
		quene = quene[end:]
		end = len(quene)
	}
	return res
}