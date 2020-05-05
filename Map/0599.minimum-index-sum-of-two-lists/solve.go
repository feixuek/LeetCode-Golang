func findRestaurant(list1 []string, list2 []string) []string {
	h1 := make(map[string]int)
	for i := 0; i < len(list1); i++ {
		h1[list1[i]] = i
	}
	h2 := make(map[int][]string)
	for i := 0; i < len(list2); i++ {
		if n, ok := h1[list2[i]]; ok {
			length := i + n
			tmp, ok := h2[length]
			if ok {
				h2[length] = append(tmp, list2[i])
			} else {
				h2[length] = []string{list2[i]}
			}
		}
	}
	res := []string{}
	min := 1<<32 - 1
	for k, v := range h2 {
		if k < min {
			min = k
			res = v
		}
	}
	return res
}