func canPartition(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%2 != 0 {
		return false
	}
	sum /= 2
	var dp [][]bool
	for i := 0; i < len(nums); i++ {
		dp = append(dp, make([]bool, sum+1))
		dp[i][0] = true
	}
	for i := 1; i < len(nums); i++ {
		for j := 1; j < sum+1; j++ {
			switch {
			case j-nums[i-1] < 0:
				dp[i][j] = dp[i-1][j]
			default:
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i-1]]
			}
		}
	}
	return dp[len(nums)-1][sum]
}