/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
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

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	tn := &TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
	if root == nil {
		return tn
	}
	pre, tmp := root, root
	for tmp != nil {
		pre = tmp
		switch {
		case tmp.Val > val:
			tmp = tmp.Left
		default:
			tmp = tmp.Right
		}
	}
	if pre.Val > val {
		pre.Left = tn
		return root
	}
	pre.Right = tn
	return root
}