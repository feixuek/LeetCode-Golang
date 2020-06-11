func new21Game(N int, K int, W int) float64 {
	if K == 0 {
		return 1.0
	}
	dp := make([]float64, K+W+1)
	for i := K; i < K+W && i <= N; i++ {
		dp[i] = 1.0
	}
	for i := 1; i <= W; i++ {
		dp[K-1] += dp[K-1+i] / float64(W)
	}
	for i := K - 2; i >= 0; i-- {
		dp[i] = dp[i+1] - (dp[i+1+W]-dp[i+1])/float64(W)
	}
	return dp[0]
}