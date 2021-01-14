func combinationSum3(k int, n int) [][]int {
	res := make([][]int, 0)
	backtrack(k, n, 1, &res, &[]int{})
	return res
}

func backtrack(k, n, start int, res *[][]int, path *[]int) {
	if n == 0 && k == 0 {
		dst := make([]int, len(*path))
		copy(dst, *path)
		*res = append(*res, dst)
		return
	}
	for i := start; i < 10; i++ {
		n -= i
		*path = append(*path, i)
		backtrack(k-1, n, i+1, res, path)
		n += i
		*path = (*path)[:len(*path)-1]
	}
}