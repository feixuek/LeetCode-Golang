func powerfulIntegers(x int, y int, bound int) []int {
	helper := make(map[int]bool)
	for i := 1; i < bound; i *= x {
		for j := 1; i+j <= bound; j *= y {
			helper[i+j] = true
			if y == 1 {
				break
			}
		}
		if x == 1 {
			break
		}
	}
	var ret []int
	for v := range helper {
		ret = append(ret, v)
	}
	return ret
}