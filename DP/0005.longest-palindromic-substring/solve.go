func longestPalindrome(s string) string {
	var res string
	for i := 0; i < len(s); i++ {
		si := palindrome(s, i, i)
		si1 := palindrome(s, i, i+1)
		if len(si) < len(si1) {
			si = si1
		}
		if len(res) < len(si) {
			res = si
		}
	}
	return res
}

func palindrome(s string, start, end int) string {
	for start >= 0 && end < len(s) && s[start] == s[end] {
		start--
		end++
	}
	return s[start+1 : end]
}