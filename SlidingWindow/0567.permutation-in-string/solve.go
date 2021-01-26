func checkInclusion(s1 string, s2 string) bool {
	source := []byte(s1)
	target := []byte(s2)
	need := make(map[byte]int)
	for _, v := range source {
		need[v]++
	}
	left, right, valid := 0, 0, 0
	window := make(map[byte]int)
	for right < len(target) {
		c := target[right]
		num, ok := need[c]
		if ok {
			window[c]++
			if window[c] == num {
				valid++
			}
		}
		right++
		for (right - left) >= len(source) {
			if len(need) == valid {
				return true
			}
			ch := target[left]
			num, ok := need[ch]
			if ok {
				if window[ch] == num {
					valid--
				}
				window[ch]--
			}
			left++
		}
	}
	return false
}
