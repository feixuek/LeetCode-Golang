func isSubPath(head *ListNode, root *TreeNode) bool {
	if root == nil {
		return false
	}
	return help(head, root) || isSubPath(head, root.Left) || isSubPath(head, root.Right)
}

func help(head *ListNode, root *TreeNode) bool {
	switch {
	case head == nil:
		return true
	case root == nil:
		return false
	default:
		return head.Val == root.Val && (help(head.Next, root.Left) || help(head.Next, root.Right))
	}
}