func lengthOfLongestSubstring(s string) int {
	window := make(map[byte]int)
	left, right, res := 0, 0, 0
	for right < len(s) {
		c := s[right]
		right++
		window[c]++
		for window[c] > 1 {
			ch := s[left]
			window[ch]--
			left++
		}
		res = max(res, right-left)
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}