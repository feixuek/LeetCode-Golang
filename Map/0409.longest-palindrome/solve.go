func longestPalindrome(s string) int {
	helper := make(map[rune]int)
	for _, v := range s {
		helper[v]++
	}
	ret := 0
	for _, n := range helper {
		ret += n / 2 * 2
		if n%2 == 1 && ret%2 == 0 {
			ret++
		}
	}
	return ret
}