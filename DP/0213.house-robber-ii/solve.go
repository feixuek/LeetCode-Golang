func rob(nums []int) int {
	switch len(nums) {
	case 0:
		return 0
	case 1:
		return nums[0]
	default:
		return max(helper(nums, 0, len(nums)-2), helper(nums, 1, len(nums)-1))
	}
}
func helper(nums []int, start, end int) int {
	n := end - start
	switch n {
	case 0:
		return nums[start]
	case 1:
		return max(nums[start], nums[end])
	default:
		dp := make([]int, end-start+1)
		dp[0] = nums[start]
		dp[1] = max(nums[start], nums[start+1])
		for i := 2; i < len(dp); i++ {
			idx := start + i
			dp[i] = max(nums[idx]+dp[i-2], dp[i-1])
		}
		return dp[len(dp)-1]
	}
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}