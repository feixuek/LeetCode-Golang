func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i < len(dp); i++ {
		dp[i] = 1 << 32
	}
	for i := 0; i < len(dp); i++ {
		for _, v := range coins {
			idx := i - v
			if idx < 0 {
				continue
			}
			dp[i] = min(dp[idx]+1, dp[i])
		}
	}
	if dp[amount] == 1<<32 {
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