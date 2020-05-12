func minDiffInBST(root *TreeNode) int {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return 0
	}
	ret := helper(root)
	left := minDiffInBST(root.Left)
	right := minDiffInBST(root.Right)
	if left != 0 && left < ret {
		ret = left
	}
	if right != 0 && right < ret {
		ret = right
	}
	return ret
}

func helper(root *TreeNode) int {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return 0
	}
	left := root.Left
	right := root.Right
	for left != nil && left.Right != nil {
		left = left.Right
	}
	for right != nil && right.Left != nil {
		right = right.Left
	}
	if left == nil {
		return right.Val - root.Val
	}
	if right == nil {
		return root.Val - left.Val
	}
	return min(root.Val-left.Val, right.Val-root.Val)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}