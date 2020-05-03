func copyRandomList(head *Node) *Node {
	if head == nil {
		return head
	}
	curr := head
	for curr != nil {
		tmp := &Node{
			Val:    curr.Val,
			Next:   curr.Next,
			Random: nil,
		}
		curr.Next = tmp
		curr = tmp.Next
	}
	randomCurr := head
	for randomCurr != nil {
		if randomCurr.Random != nil {
			randomCurr.Next.Random = randomCurr.Random.Next
		}
		randomCurr = randomCurr.Next.Next
	}
	res := new(Node)
	headHead := head
	resHead := res
	for headHead != nil {
		tmp := headHead.Next
		resHead.Next = tmp
		resHead = tmp
		headHead.Next = tmp.Next
		headHead = tmp.Next
	}
	return res.Next
}