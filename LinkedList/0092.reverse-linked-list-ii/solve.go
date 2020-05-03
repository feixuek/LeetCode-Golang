func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil {
		return head
	}
	res := &ListNode{
		Val:  0,
		Next: head,
	}
	pre := res
	for i := 0; i < m-1; i++ {
		pre = pre.Next
	}
	start := pre.Next
	then := start.Next
	for i := 0; i < n-m; i++ {
		start.Next = then.Next
		then.Next = pre.Next
		pre.Next = then
		then = start.Next
	}
	return res.Next
}