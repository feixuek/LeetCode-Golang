type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
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

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	queue, helper := new(Queue), new(Queue)
	queue.Push(root)
	for !queue.Empty() || !helper.Empty() {
		switch {
		case !queue.Empty():
			for !queue.Empty() {
				helper.Push(queue.Pop())
			}
		case !helper.Empty():
			var pre *Node
			for !helper.Empty() {
				elem := helper.Pop()
				if pre != nil {
					pre.Next = elem
				}
				if elem.Left != nil {
					queue.Push(elem.Left)
				}
				if elem.Right != nil {
					queue.Push(elem.Right)
				}
				pre = elem
			}
		}
	}
	return root
}
