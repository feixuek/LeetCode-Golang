func partition(head *ListNode, x int) *ListNode {
	small := new(ListNode)
	big := new(ListNode)
	smallHead := small
	bigHead := big
	for head != nil {
		if head.Val < x {
			smallHead.Next = head
			smallHead = smallHead.Next
		} else {
			bigHead.Next = head
			bigHead = bigHead.Next
		}
		head = head.Next
	}
	bigHead.Next = nil
	smallHead.Next = big.Next
	return small.Next
}