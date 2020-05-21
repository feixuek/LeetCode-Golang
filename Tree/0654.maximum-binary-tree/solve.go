func constructMaximumBinaryTree(nums []int) *TreeNode {
	return helper(nums, 0, len(nums)-1)
}

func helper(nums []int, start, end int) *TreeNode {
	if start > end {
		return nil
	}
	max, idx := nums[start], start
	for i := start + 1; i <= end; i++ {
		if nums[i] > max {
			max = nums[i]
			idx = i
		}
	}
	root := &TreeNode{
		Val:   max,
		Left:  helper(nums, start, idx-1),
		Right: helper(nums, idx+1, end),
	}
	return root
}