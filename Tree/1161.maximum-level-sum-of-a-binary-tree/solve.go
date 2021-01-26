func maxLevelSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	quene := []*TreeNode{root}
	end := len(quene)
	res := 1
	level := 1
	sum := -(1 << 31)
	for end != 0 {
		tmp := 0
		for i := 0; i < end; i++ {
			elem := quene[i]
			tmp += elem.Val
			if elem.Left != nil {
				quene = append(quene, elem.Left)
			}
			if elem.Right != nil {
				quene = append(quene, elem.Right)
			}
		}
		if tmp > sum {
			sum = tmp
			res = level
		}
		quene = quene[end:]
		end = len(quene)
		level++
	}
	return res
}