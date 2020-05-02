	if head == nil {
		return head
	}
	ret, tail := head, head
	length := 1
	for tail.Next != nil {
		tail = tail.Next
		length++
	}
	tail.Next = head
	k %= length
	if k > 0 {
		for i := 0; i < length-k; i++ {
			tail = tail.Next
		}
	}
	ret = tail.Next
	tail.Next = nil
	return ret
}