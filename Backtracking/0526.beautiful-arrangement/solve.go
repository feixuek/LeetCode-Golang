func countArrangement(n int) int {
	res := new(int)
	filter := make(map[int]bool)
	backtrack(n, 0, &filter, res)
	return *res
}

func backtrack(n, index int, filter *map[int]bool, res *int) {
	if index == n {
		(*res)++
		return
	}

	for i := 1; i <= n; i++ {
		if (*filter)[i] {
			continue
		}
		target := index + 1
		if target%i == 0 || i%target == 0 {
			(*filter)[i] = true
			backtrack(n, index+1, filter, res)
			(*filter)[i] = false
		}
	}
}