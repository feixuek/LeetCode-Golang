func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	odd, oddHead := head, head
	even, evenHead := head.Next, head.Next
	for evenHead != nil && evenHead.Next != nil {
		tmp := evenHead.Next
		oddHead.Next = tmp
		oddHead = tmp
		evenHead.Next = tmp.Next
		evenHead = tmp.Next
	}
	oddHead.Next = even
	return odd
}