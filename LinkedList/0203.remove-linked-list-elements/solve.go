func removeElements(head *ListNode, val int) *ListNode {
	root := &ListNode{
		Val:  0,
		Next: head,
	}
	curr := root
	for curr.Next != nil {
		if curr.Next.Val == val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}
	return root.Next
}