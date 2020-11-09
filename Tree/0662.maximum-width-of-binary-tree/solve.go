type SloutionNode struct {
	Node  *TreeNode
	Index int
}

type Queue struct {
	data []*SloutionNode
}

func (q *Queue) Enqueue(elem *SloutionNode) {
	q.data = append(q.data, elem)
}

func (q *Queue) Dequeue() *SloutionNode {
	if q.Empty() {
		return nil
	}
	elem := q.data[0]
	q.data = q.data[1:]
	return elem
}

func (q *Queue) Empty() bool {
	return len(q.data) == 0
}

func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	q := new(Queue)
	help := new(Queue)
	ret := 0
	q.Enqueue(&SloutionNode{Node: root, Index: 0})
	for !q.Empty() || !help.Empty() {
		switch {
		case !q.Empty():
			var left, right *SloutionNode
			for !q.Empty() {
				elem := q.Dequeue()
				if left == nil {
					left = elem
				}
				if elem.Node.Left != nil {
					help.Enqueue(&SloutionNode{Node: elem.Node.Left, Index: elem.Index * 2})
				}
				if elem.Node.Right != nil {
					help.Enqueue(&SloutionNode{Node: elem.Node.Right, Index: elem.Index*2 + 1})
				}
				if q.Empty() {
					right = elem
				}
			}
			length := right.Index - left.Index + 1
			if length > ret {
				ret = length
			}
		case !help.Empty():
			for !help.Empty() {
				q.Enqueue(help.Dequeue())
			}
		}
	}
	return ret
}