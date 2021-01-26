func rob(nums []int) int {
	switch len(nums) {
	case 0:
		return 0
	case 1:
		return nums[0]
	default:
		dp := make([]int, len(nums))
		dp[0] = nums[0]
		dp[1] = max(nums[0], nums[1])
		for i := 2; i < len(nums); i++ {
			dp[i] = max(dp[i-1], dp[i-2]+nums[i])
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