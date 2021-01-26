type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	quene := []*Node{root}
	end := len(quene)
	for end != 0 {
		var prev *Node
		for i := 0; i < end; i++ {
			elem := quene[i]
			if prev != nil {
				prev.Next = elem
			}
			if elem.Left != nil {
				quene = append(quene, elem.Left)
			}
			if elem.Right != nil {
				quene = append(quene, elem.Right)
			}
			prev = elem
		}
		quene = quene[end:]
		if prev != nil {
			prev.Next = nil
		}
		end = len(quene)
	}
	return root
}