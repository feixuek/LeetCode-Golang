type Queue struct {
	Data []*TreeNode
}

func (q *Queue) Empty() bool {
	return len(q.Data) == 0
}

func (q *Queue) Push(node *TreeNode) {
	q.Data = append(q.Data, node)
}

func (q *Queue) Pop() (node *TreeNode) {
	if !q.Empty() {
		node = q.Data[0]
		q.Data = q.Data[1:]
		return
	}
	return nil
}

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}
	queue, helper := new(Queue), new(Queue)
	queue.Push(root)
	var res []float64
	for !queue.Empty() || !helper.Empty() {
		if !queue.Empty() {
			for !queue.Empty() {
				helper.Push(queue.Pop())
			}
		} else {
			num := 0
			sum := 0
			for !helper.Empty() {
				elem := helper.Pop()
				num++
				sum += elem.Val
				if elem.Left != nil {
					queue.Push(elem.Left)
				}
				if elem.Right != nil {
					queue.Push(elem.Right)
				}
			}
			if num > 0 {
				res = append(res, float64(sum)/float64(num))
			}
		}
	}
	return res
}
