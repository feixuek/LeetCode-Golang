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

func deepestLeavesSum(root *TreeNode) int {
	q := new(Queue)
	q.Enqueue(root)
	sum := 0
	for !q.Empty() {
		length := q.Len()
		if length != 0 {
			sum = 0
		}
		for i := 0; i < length; i++ {
			elem := q.Dequeue()
			sum += elem.Val
			if elem.Left != nil {
				q.Enqueue(elem.Left)
			}
			if elem.Right != nil {
				q.Enqueue(elem.Right)
			}
		}
	}
	return sum
}