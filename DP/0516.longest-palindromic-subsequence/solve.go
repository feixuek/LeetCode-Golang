func longestPalindromeSubseq(s string) int {
	length := len(s)
	var res [][]int
	for i := 0; i < length; i++ {
		res = append(res, make([]int, length))
		res[i][i] = 1
	}
	for i := length - 1; i >= 0; i-- {
		for j := i + 1; j < length; j++ {
			if s[i] == s[j] {
				res[i][j] = res[i+1][j-1] + 2
			} else {
				res[i][j] = max(res[i+1][j], res[i][j-1])
			}
		}
	}
	return res[0][length-1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}