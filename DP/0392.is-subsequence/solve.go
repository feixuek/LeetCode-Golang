func isSubsequence(s string, t string) bool {
	i, j := 0, 0
	for j < len(t) && i < len(s) {
		switch {
		case s[i] == t[j]:
			i++
			j++
		default:
			j++
		}
	}
	return i == len(s)
}