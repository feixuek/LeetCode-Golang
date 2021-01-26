func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	max := 2 << 32
	for i := 1; i < len(dp); i++ {
		dp[i] = max
	}
	dp[0] = 0
	for i := 1; i < len(dp); i++ {
		for j := 0; j < len(coins); j++ {
			idx := i - coins[j]
			if idx < 0 {
				continue
			}
			dp[i] = min(dp[i], dp[idx]+1)
		}
	}
	if dp[amount] == max {
		return -1
	}
	return dp[amount]
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}