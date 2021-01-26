func isEvenOddTree(root *TreeNode) bool {
	quene := []*TreeNode{root}
	level := 0
	end := len(quene)
	for end != 0 {
		var pre *TreeNode
		for i := 0; i < end; i++ {
			elem := quene[i]
			if level%2 == 0 {
				if elem.Val%2 == 0 {
					return false
				}
				if pre != nil && (elem.Val-pre.Val) <= 0 {
					return false
				}
			} else {
				if elem.Val%2 != 0 {
					return false
				}
				if pre != nil && (elem.Val-pre.Val) >= 0 {
					return false
				}
			}
			pre = elem
			if elem.Left != nil {
				quene = append(quene, elem.Left)
			}
			if elem.Right != nil {
				quene = append(quene, elem.Right)
			}
		}
		quene = quene[end:]
		end = len(quene)
		level++
	}
	return true
}