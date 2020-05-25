/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
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

func kthSmallest(root *TreeNode, k int) int {
	if root == nil {
		return -1
	}
	stack := new(Stack)
	for root != nil || !stack.Empty() {
		switch {
		case root != nil:
			stack.Push(root)
			root = root.Left
		default:
			elem := stack.Pop()
			k--
			if k == 0 {
				return elem.Val
			}
			if elem.Right != nil {
				root = elem.Right
			}
		}
	}
	return -1
}