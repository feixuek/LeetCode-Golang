var (
	data = []byte{'a', 'b', 'c'}
)

func getHappyString(n int, k int) string {
	res := ""
	backtrack(n, &k, &[]byte{}, &res)
	return res
}

func backtrack(n int, k *int, path *[]byte, res *string) {
	if n == len(*path) {
		(*k)--
		if *k == 0 {
			*res = string(*path)
		}
		return
	}
	for i := 0; i < 3; i++ {
		if len(*path) > 0 && (*path)[len(*path)-1] == data[i] {
			continue
		}
		*path = append(*path, data[i])
		backtrack(n, k, path, res)
		*path = (*path)[:len(*path)-1]
	}
}