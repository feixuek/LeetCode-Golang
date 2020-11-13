type FindElements struct {
	Data map[int]bool
}

func Constructor(root *TreeNode) FindElements {
	elem := &FindElements{
		Data: make(map[int]bool),
	}
	if root != nil {
		root.Val = 0
	}
	help(root, elem)
	return *elem
}
func help(root *TreeNode, elem *FindElements) {
	if root == nil {
		return
	}
	elem.Data[root.Val] = true
	if root.Left != nil {
		root.Left.Val = root.Val*2 + 1
	}
	if root.Right != nil {
		root.Right.Val = root.Val*2 + 2
	}
	help(root.Left, elem)
	help(root.Right, elem)
}

func (this *FindElements) Find(target int) bool {
	if this == nil || this.Data == nil {
		return false
	}
	if _, ok := this.Data[target]; ok {
		return true
	}
	return false
}