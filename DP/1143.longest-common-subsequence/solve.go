func longestCommonSubsequence(text1 string, text2 string) int {
	var dp [][]int
	for i := 0; i <= len(text1); i++ {
		dp = append(dp, make([]int, len(text2)+1))
	}
	for i := 1; i <= len(text1); i++ {
		for j := 1; j <= len(text2); j++ {
			switch {
			case text1[i-1] == text2[j-1]:
				dp[i][j] = dp[i-1][j-1] + 1
			default:
				dp[i][j] = max(dp[i-1][j], max(dp[i][j-1], dp[i-1][j-1]))
			}
		}
	}
	return dp[len(text1)][len(text2)]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}