func numJewelsInStones(J string, S string) int {
	helper := make(map[rune]int)
	for _, v := range J {
		helper[v] = 0
	}
	res := 0
	for _, v := range S {
		if _, ok := helper[v]; ok {
			res++
		}
	}
	return res
}