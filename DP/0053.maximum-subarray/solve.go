func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		sum := nums[i] + dp[i-1]
		if nums[i] > sum {
			dp[i] = nums[i]
		} else {
			dp[i] = sum
		}
	}
	res := dp[0]
	for i := 1; i < len(dp); i++ {
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}