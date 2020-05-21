type Node struct {
	Val      int
	Children []*Node
}

type Queue struct {
	Data []*Node
}

func (q *Queue) Push(elem *Node) {
	q.Data = append(q.Data, elem)
}

func (q *Queue) Pop() (elem *Node) {
	if q.Empty() {
		return
	}
	elem = q.Data[0]
	q.Data = q.Data[1:]
	return
}

func (q *Queue) Peek() (elem *Node) {
	if q.Empty() {
		return
	}
	elem = q.Data[0]
	return
}

func (q *Queue) Empty() bool {
	return len(q.Data) == 0
}

func levelOrder(root *Node) [][]int {
	if root == nil {
		return nil
	}
	queue, helper := new(Queue), new(Queue)
	queue.Push(root)
	var res [][]int
	for !queue.Empty() || !helper.Empty() {
		switch {
		case !queue.Empty():
			for !queue.Empty() {
				helper.Push(queue.Pop())
			}
		default:
			var tmp []int
			for !helper.Empty() {
				elem := helper.Pop()
				tmp = append(tmp, elem.Val)
				for _, v := range elem.Children {
					queue.Push(v)
				}
			}
			res = append(res, tmp)
		}
	}
	return res
}