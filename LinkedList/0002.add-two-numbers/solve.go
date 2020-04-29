func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sum := 0
	ret := new(ListNode)
	curr := ret
	for l1 != nil || l2 != nil {
		sum /= 10
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		curr.Next = &ListNode{
			Val:  sum % 10,
			Next: nil,
		}
		curr = curr.Next
	}
	if sum/10 == 1 {
		curr.Next = &ListNode{
			Val:  1,
			Next: nil,
		}
	}
	return ret.Next
}