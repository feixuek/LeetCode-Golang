type Queue struct {
	data []*TreeNode
}

func (q *Queue) Enqueue(elem *TreeNode) {
	q.data = append(q.data, elem)
}

func (q *Queue) Dequeue() *TreeNode {
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

func (q *Queue) Len() int {
	return len(q.data)
}

func isCompleteTree(root *TreeNode) bool {
	q := new(Queue)
	q.Enqueue(root)
	flag := false
	for !q.Empty() {
		size := q.Len()
		for i := 0; i < size; i++ {
			elem := q.Dequeue()
			if elem != nil {
				if flag {
					return false
				}
				q.Enqueue(elem.Left)
				q.Enqueue(elem.Right)
			} else {
				flag = true
			}
		}
	}
	return true
}