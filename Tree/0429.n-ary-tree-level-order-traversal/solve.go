func levelOrder(root *Node) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0)
	quene := []*Node{root}
	end := len(quene)
	for end != 0 {
		data := []int{}
		for i := 0; i < end; i++ {
			elem := quene[i]
			data = append(data, elem.Val)
			for j := 0; j < len(elem.Children); j++ {
				if elem.Children[j] != nil {
					quene = append(quene, elem.Children[j])
				}
			}
		}
		res = append(res, data)
		quene = quene[end:]
		end = len(quene)
	}
	return res
}