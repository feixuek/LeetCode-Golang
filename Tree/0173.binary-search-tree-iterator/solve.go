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

func find(root *TreeNode, val int) *TreeNode {
	for root != nil {
		switch {
		case root.Val == val:
			return root
		case root.Val > val:
			root = root.Left
		default:
			root = root.Right
		}
	}
	return nil
}

type BSTIterator struct {
	elem   *TreeNode
	helper *Stack
}

func Constructor(root *TreeNode) BSTIterator {
	ret := BSTIterator{elem: root, helper: new(Stack)}
	for root != nil {
		ret.helper.Push(root)
		root = root.Left
	}
	return ret
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	elem := this.helper.Pop()
	if elem == nil {
		return -1
	}
	ret := elem.Val
	exist := find(this.elem, ret)
	if exist != nil {
		tmp := exist.Right
		for tmp != nil {
			this.helper.Push(tmp)
			tmp = tmp.Left
		}
	}
	return ret
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return !this.helper.Empty()
}