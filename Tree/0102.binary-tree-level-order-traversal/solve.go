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

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	queue, helper := new(Queue), new(Queue)
	var res [][]int
	queue.Push(root)
	for !queue.Empty() || !helper.Empty() {
		if !queue.Empty() {
			for !queue.Empty() {
				helper.Push(queue.Pop())
			}
		} else {
			tmp := []int{}
			for !helper.Empty() {
				elem := helper.Pop()
				tmp = append(tmp, elem.Val)
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