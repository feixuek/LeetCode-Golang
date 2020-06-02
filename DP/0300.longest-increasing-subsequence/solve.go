func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	ret := 0
	for i := 0; i < len(dp); i++ {
		if dp[i] > ret {
			ret = dp[i]
		}
	}
	return ret
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}