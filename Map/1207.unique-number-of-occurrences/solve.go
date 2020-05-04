func uniqueOccurrences(arr []int) bool {
	helper := make(map[int]int)
	for _, v := range arr {
		n, ok := helper[v]
		if ok {
			helper[v] = n + 1
		} else {
			helper[v] = 1
		}
	}
	h2 := make(map[int]bool)
	for _, v := range helper {
		if ok := h2[v]; ok {
			return false
		}
		h2[v] = true
	}
	return true
}