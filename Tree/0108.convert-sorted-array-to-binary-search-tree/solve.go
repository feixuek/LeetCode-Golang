func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	return helper(nums, 0, len(nums)-1)
}

func helper(nums []int, low, high int) *TreeNode {
	if low > high {
		return nil
	}
	mid := (low + high) / 2
	tree := &TreeNode{
		Val: nums[mid],
	}
	tree.Left = helper(nums, low, mid-1)
	tree.Right = helper(nums, mid+1, high)
	return tree
}