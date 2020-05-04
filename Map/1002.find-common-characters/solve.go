func commonChars(A []string) []string {
	if len(A) == 0 {
		return nil
	}
	helper := make(map[rune]int)
	for _, v := range A[0] {
		helper[v]++
	}
	for i := 1; i < len(A); i++ {
		tmp := make(map[rune]int)
		for _, v := range A[i] {
			n, ok := helper[v]
			if ok && n > 0 {
				tmp[v]++
				helper[v]--
			}
		}
		helper = tmp
	}
	res := []string{}
	for v, n := range helper {
		for n > 0 {
			res = append(res, string(v))
			n--
		}
	}
	return res
}