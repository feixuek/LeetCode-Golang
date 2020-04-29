func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	ret := new(ListNode)
	now := ret
	h1 := l1
	h2 := l2
	for h1 != nil && h2 != nil {
		if h1.Val <= h2.Val {
			now.Next = h1
			h1 = h1.Next
		} else {
			now.Next = h2
			h2 = h2.Next
		}
		now = now.Next
	}
	if h1 != nil {
		now.Next = h1
	}
	if h2 != nil {
		now.Next = h2
	}
	return ret.Next
}