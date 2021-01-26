func isCompleteTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	quene := []*TreeNode{root}
	end := len(quene)
	lastRow := false
	for end != 0 {
		for i := 0; i < end; i++ {
			elem := quene[i]
			if elem != nil {
				if lastRow {
					return false
				}
				quene = append(quene, elem.Left)
				quene = append(quene, elem.Right)
			} else {
				lastRow = true
			}
		}
		quene = quene[end:]
		end = len(quene)
	}
	return true
}