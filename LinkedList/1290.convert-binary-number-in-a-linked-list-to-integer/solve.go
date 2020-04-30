func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	for head != nil {
		tmp := head.Next
		head.Next = pre
		pre = head
		head = tmp
	}
	return pre
}

func getDecimalValue(head *ListNode) int {
	reverse := reverseList(head)
	sum := 0
	exp := 1
	for reverse != nil {
		sum += reverse.Val * exp
		exp *= 2
		reverse = reverse.Next
	}
	return sum
}