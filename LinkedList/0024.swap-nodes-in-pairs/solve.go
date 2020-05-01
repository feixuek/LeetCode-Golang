func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	ret := head.Next
	head.Next = swapPairs(head.Next.Next)
	ret.Next = head
	return ret
}