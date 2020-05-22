/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	tn := &TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
	if root == nil {
		return tn
	}
	curr := root
	for {
		switch {
		case curr.Val > val:
			switch {
			case curr.Left == nil:
				curr.Left = tn
				goto RET
			default:
				curr = curr.Left
			}
		default:
			switch {
			case curr.Right == nil:
				curr.Right = tn
				goto RET
			default:
				curr = curr.Right
			}
		}
	}
RET:
	return root
}