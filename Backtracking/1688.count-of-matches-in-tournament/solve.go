func numberOfMatches(n int) int {
	res := new(int)
	helper(n, res)
	return *res
}

func helper(n int, res *int) {
	if n <= 1 {
		return
	}
	if n == 2 {
		*res += 1
		return
	}
	num := n / 2
	*res += num
	helper(num+n%2, res)
}