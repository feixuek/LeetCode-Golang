func longestPalindromeSubseq(s string) int {
	if len(s) == 0 {
		return 0
	}
	var dp [][]int
	for i := 0; i < len(s); i++ {
		dp = append(dp, make([]int, len(s)))
		dp[i][i] = 1
	}
	for i := len(s) - 1; i >= 0; i-- {
		for j := i + 1; j < len(s); j++ {
			switch {
			case s[i] == s[j]:
				dp[i][j] = dp[i+1][j-1] + 2
			default:
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][len(s)-1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}