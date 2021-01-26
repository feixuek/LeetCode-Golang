func minWindow(s string, t string) string {
	source := []byte(t)
	target := []byte(s)
	need := make(map[byte]int)
	for _, v := range source {
		need[v]++
	}
	left, right, start, valid := 0, 0, 0, 0
	length := 1 << 32
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
		for valid == len(need) {
			if right-left < length {
				length = right - left
				start = left
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
	if length == 1<<32 {
		return ""
	}
	return string(target[start : start+length])
}