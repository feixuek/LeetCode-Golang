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

func largestValues(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	queue, helper := new(Queue), new(Queue)
	queue.Push(root)
	var res []int
	for !queue.Empty() || !helper.Empty() {
		if !queue.Empty() {
			helper.Push(queue.Pop())
		} else {
			tmp := -1 << 31
			for !helper.Empty() {
				elem := helper.Pop()
				if elem.Val > tmp {
					tmp = elem.Val
				}
				if elem.Left != nil {
					queue.Push(elem.Left)
				}
				if elem.Right != nil {
					queue.Push(elem.Right)
				}
			}
			res = append(res, tmp)
		}
	}
	return res
}