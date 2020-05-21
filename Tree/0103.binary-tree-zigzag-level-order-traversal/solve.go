type Stack struct {
	Data []*TreeNode
}

func (s *Stack) Push(elem *TreeNode) {
	s.Data = append(s.Data, elem)
}

func (s *Stack) Pop() (elem *TreeNode) {
	if s.Empty() {
		return
	}
	elem = s.Data[len(s.Data)-1]
	s.Data = s.Data[:len(s.Data)-1]
	return
}

func (s *Stack) Top() (elem *TreeNode) {
	if s.Empty() {
		return
	}
	elem = s.Data[len(s.Data)-1]
	return
}

func (s *Stack) Empty() bool {
	return len(s.Data) == 0
}

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

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	queue, helper := new(Queue), new(Queue)
	stack := new(Stack)
	queue.Push(root)
	var res [][]int
	level := 1
	for !queue.Empty() || !helper.Empty() {
		if !queue.Empty() {
			for !queue.Empty() {
				helper.Push(queue.Pop())
			}
		} else {
			tmp := []int{}
			for !helper.Empty() {
				elem := helper.Pop()
				if level%2 == 0 {
					stack.Push(elem)
				} else {
					tmp = append(tmp, elem.Val)
				}
				if elem.Left != nil {
					queue.Push(elem.Left)
				}
				if elem.Right != nil {
					queue.Push(elem.Right)
				}
			}
			level++
			for !stack.Empty() {
				tmp = append(tmp, stack.Pop().Val)
			}
			res = append(res, tmp)
		}
	}
	return res
}