func isAnagram(s string, t string) bool {
	helper := make(map[rune]int)
	for _, v := range s {
		helper[v]++
	}
	for _, v := range t {
		count, ok := helper[v]
		if !ok {
			return false
		}
		count -= 1
		if count == 0 {
			delete(helper, v)
		} else {
			helper[v] = count
		}
	}
	if len(helper) != 0 {
		return false
	}
	return true
}