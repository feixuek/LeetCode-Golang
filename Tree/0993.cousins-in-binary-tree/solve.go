type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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

func isCousins(root *TreeNode, x int, y int) bool {
	queue, helper := new(Queue), new(Queue)
	queue.Push(root)
	for !queue.Empty() || !helper.Empty() {
		if !queue.Empty() {
			for !queue.Empty() {
				helper.Push(queue.Pop())
			}
		} else {
			flagX, flagY := false, false
			for !helper.Empty() {
				elem := helper.Pop()
				if elem.Left != nil && elem.Right != nil {
					if (elem.Left.Val == x && elem.Right.Val == y) || (elem.Left.Val == y && elem.Right.Val == x) {
						return false
					}
				}
				if elem.Left != nil {
					switch elem.Left.Val {
					case x:
						flagX = true
					case y:
						flagY = true
					}
					queue.Push(elem.Left)
				}
				if elem.Right != nil {
					switch elem.Right.Val {
					case x:
						flagX = true
					case y:
						flagY = true
					}
					queue.Push(elem.Right)
				}
			}
			if flagX && flagY {
				return true
			}
		}
	}
	return false
}