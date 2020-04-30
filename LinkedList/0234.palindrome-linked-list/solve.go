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

func isPalindrome(head *ListNode) bool {
	slow, fast := head, head
	for ; fast != nil && fast.Next != nil; fast = fast.Next.Next {
		slow = slow.Next
	}
	reverse := reverseList(slow)
	for head != nil && reverse != nil {
		if head.Val != reverse.Val {
			return false
		}
		head = head.Next
		reverse = reverse.Next
	}
	return true
}