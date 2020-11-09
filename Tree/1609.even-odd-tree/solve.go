type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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

func isEvenOddTree(root *TreeNode) bool {
	q := new(Queue)
	help := new(Queue)
	q.Enqueue(root)
	level := 1
	for !q.Empty() || !help.Empty() {
		switch {
		case !q.Empty():
			var pre *TreeNode
			for !q.Empty() {
				elem := q.Dequeue()
				if level%2 == 0 {
					if elem.Val%2 != 0 {
						return false
					}
					if pre != nil && elem.Val >= pre.Val {
						return false
					}
				} else {
					if elem.Val%2 == 0 {
						return false
					}
					if pre != nil && elem.Val <= pre.Val {
						return false
					}
				}
				if elem.Left != nil {
					help.Enqueue(elem.Left)
				}
				if elem.Right != nil {
					help.Enqueue(elem.Right)
				}
				pre = elem
			}
		case !help.Empty():
			for !help.Empty() {
				q.Enqueue(help.Dequeue())
			}
			level++
		}
	}
	return true
}