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

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	queue, helper := new(Queue), new(Queue)
	queue.Push(root)
	var res [][]int
	for !queue.Empty() || !helper.Empty() {
		if !queue.Empty() {
			for !queue.Empty() {
				helper.Push(queue.Pop())
			}
		} else {
			tmp := make([]int, 0)
			for !helper.Empty() {
				node := helper.Pop()
				tmp = append(tmp, node.Val)
				if node.Left != nil {
					queue.Push(node.Left)
				}
				if node.Right != nil {
					queue.Push(node.Right)
				}
			}
			res = append(res, tmp)
		}
	}
	start := 0
	end := len(res) - 1
	for start < end {
		res[start], res[end] = res[end], res[start]
		start++
		end--
	}
	return res
}