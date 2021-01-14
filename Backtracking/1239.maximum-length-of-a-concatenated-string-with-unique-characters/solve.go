func maxLength(arr []string) int {
	res := new(int)
	backtracek(arr, 0, res, &[]byte{})
	return *res
}

func backtracek(arr []string, start int, res *int, path *[]byte) {
	ok := true
	filter := make(map[byte]bool)
	for i := 0; i < len(*path); i++ {
		if exist := filter[(*path)[i]]; exist {
			ok = false
			break
		}
		filter[(*path)[i]] = true
	}
	if ok && len(*path) > *res {
		*res = len(*path)
	}
	for i := start; i < len(arr); i++ {
		data := []byte(arr[i])
		*path = append(*path, data...)
		backtracek(arr, i+1, res, path)
		*path = (*path)[:len(*path)-len(data)]
	}
}