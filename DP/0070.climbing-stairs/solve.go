func climbStairs(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	res := 0
	step1 := 1
	step2 := 2
	for n >= 3 {
		res = step1 + step2
		step1 = step2
		step2 = res
		n--
	}
	return res
}