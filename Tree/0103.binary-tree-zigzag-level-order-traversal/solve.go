func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0)
	quene := []*TreeNode{root}
	end := len(quene)
	level := 1
	for end != 0 {
		data := []int{}
		if level%2 == 0 {
			for i := end - 1; i >= 0; i-- {
				data = append(data, quene[i].Val)
			}
			for i := 0; i < end; i++ {
				elem := quene[i]
				if elem.Left != nil {
					quene = append(quene, elem.Left)
				}
				if elem.Right != nil {
					quene = append(quene, elem.Right)
				}
			}
		} else {
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
		}
		res = append(res, data)
		quene = quene[end:]
		end = len(quene)
		level++
	}
	return res
}