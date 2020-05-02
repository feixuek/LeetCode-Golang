func lengthOfLongestSubstring(s string) int {
	helper := map[byte]int{}
	i, j, ret := 0, 0, 0
	for ; i < len(s); i++ {
		if pos, ok := helper[s[i]]; ok {
			j = max(j, pos+1)
		}
		helper[s[i]] = i
		ret = max(ret, i-j+1)
	}
	return ret
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}