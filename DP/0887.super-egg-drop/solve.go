func superEggDrop(K int, N int) int {
	var dp [][]int
	for i := 0; i <= K; i++ {
		dp = append(dp, make([]int, N+1))
	}
	step := 0
	for dp[K][step] < N {
		step++
		for i := 1; i <= K; i++ {
			dp[i][step] = dp[i-1][step-1] + dp[i][step-1] + 1
		}
	}
	return step
}