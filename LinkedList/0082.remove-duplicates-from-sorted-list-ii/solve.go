func deleteDuplicates(head *ListNode) *ListNode {
	ret := new(ListNode)
	ret.Next = head
	pre := ret
	curr := head
	for curr != nil {
		for curr.Next != nil && curr.Val == curr.Next.Val {
			curr = curr.Next
		}
		if pre.Next == curr {
			pre = pre.Next
		} else {
			pre.Next = curr.Next
		}
		curr = curr.Next
	}
	return ret.Next
}