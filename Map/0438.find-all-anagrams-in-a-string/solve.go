func findAnagrams(s string, p string) []int {
	source := []byte(p)
	target := []byte(s)
	need := make(map[byte]int)
	for _, v := range source {
		need[v]++
	}
	left, right, valid := 0, 0, 0
	res := []int{}
	window := make(map[byte]int)
	for right < len(target) {
		c := target[right]
		right++
		num, ok := need[c]
		if ok {
			window[c]++
			if window[c] == num {
				valid++
			}
		}
		for right-left >= len(source) {
			if valid == len(need) {
				res = append(res, left)
			}
			ch := target[left]
			num, ok := need[ch]
			if ok {
				if num == window[ch] {
					valid--
				}
				window[ch]--
			}
			left++
		}
	}
	return res
}