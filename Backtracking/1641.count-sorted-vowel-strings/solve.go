var (
	data = []byte{'a', 'e', 'i', 'o', 'u'}
)

func countVowelStrings(n int) int {
	res := new(int)
	path := make([]byte, 0)
	backtrack(n, res, &path)
	return *res
}

func backtrack(n int, res *int, path *[]byte) {
	if len(*path) == n {
		(*res)++
		return
	}
	for i := 0; i < 5; i++ {
		if len(*path) > 0 && (*path)[len(*path)-1] > data[i] {
			continue
		}
		*path = append(*path, data[i])
		backtrack(n, res, path)
		(*path) = (*path)[:len(*path)-1]
	}
}