type Queue struct {
	Data []*TreeNode
}

func (q *Queue) Push(elem *TreeNode) {
	q.Data = append(q.Data, elem)
}

func (q *Queue) Pop() (elem *TreeNode) {
	if q.Empty() {
		return
	}
	elem = q.Data[0]
	q.Data = q.Data[1:]
	return
}

func (q *Queue) Peek() (elem *TreeNode) {
	if q.Empty() {
		return
	}
	elem = q.Data[0]
	return
}

func (q *Queue) Empty() bool {
	return len(q.Data) == 0
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue, helper := new(Queue), new(Queue)
	queue.Push(root)
	level, count := 0, 0
	for !queue.Empty() || !helper.Empty() {
		switch {
		case !queue.Empty():
			for !queue.Empty() {
				helper.Push(queue.Pop())
			}
		default:
			level++
			count = 0
			for !helper.Empty() {
				elem := helper.Pop()
				count++
				if elem.Left != nil {
					queue.Push(elem.Left)
				}
				if elem.Right != nil {
					queue.Push(elem.Right)
				}
			}
		}
	}
	res := 1
	for i := 1; i < level; i++ {
		res *= 2
	}
	return res + count - 1
}