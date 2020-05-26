/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	switch {
	case root.Val > key:
		root.Left = deleteNode(root.Left, key)
	case root.Val < key:
		root.Right = deleteNode(root.Right, key)
	default:
		switch {
		case root.Left == nil:
			return root.Right
		case root.Right == nil:
			return root.Left
		default:
			min := minNode(root.Right)
			root.Val = min.Val
			root.Right = deleteNode(root.Right, root.Val)
		}
	}
	return root
}

func minNode(root *TreeNode) *TreeNode {
	for root.Left != nil {
		root = root.Left
	}
	return root
}