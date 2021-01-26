func isSubsequence(s string, t string) bool {
	start := 0
	for i := 0; i < len(t); i++ {
		if start >= len(s) {
			return true
		}
		if t[i] == s[start] {
			start++
		}
	}
	return start == len(s)
}