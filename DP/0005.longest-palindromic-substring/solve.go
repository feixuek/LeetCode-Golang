func longestPalindrome(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		s1 := helper(s, i, i)
		s2 := helper(s, i, i+1)
		if len(s1) > len(res) {
			res = s1
		}
		if len(s2) > len(res) {
			res = s2
		}
	}
	return res
}

func helper(s string, start, end int) string {
	for start >= 0 && end < len(s) && s[start] == s[end] {
		start--
		end++
	}
	return s[start+1 : end]
}