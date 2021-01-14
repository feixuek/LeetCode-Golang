func maxUniqueSplit(s string) int {
	path := make([]string, 0)
	res := new(int)
	filter := make(map[string]bool)
	backtrack([]byte(s), 0, &path, res, &filter)
	return *res
}

func backtrack(data []byte, index int, path *[]string, res *int, filter *map[string]bool) {
	if len(data) == index {
		if len(*path) > *res {
			*res = len(*path)
		}
	}
	for i := index; i < len(data); i++ {
		elem := string(data[index : i+1])
		if ok := (*filter)[elem]; ok {
			continue
		}
		(*filter)[elem] = true
		*path = append(*path, elem)
		backtrack(data, i+1, path, res, filter)
		*path = (*path)[:len(*path)-1]
		delete(*filter, elem)
	}
}