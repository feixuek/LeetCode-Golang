func numTilePossibilities(tiles string) int {
	filter := make(map[string]bool)
	res := new(int)
	used := make(map[int]bool)
	for i := 0; i < len(tiles); i++ {
		used[i] = false
	}
	for i := 0; i < len(tiles); i++ {
		backtrack(i+1, []byte(tiles), &[]byte{}, &filter, used, res)
	}
	return *res
}

func backtrack(number int, tiles []byte, path *[]byte, filter *map[string]bool, used map[int]bool, res *int) {
	if number == len(*path) {
		elem := string(*path)
		if ok := (*filter)[elem]; ok {
			return
		}
		(*filter)[elem] = true
		(*res)++
		return
	}
	for i := 0; i < len(tiles); i++ {
		if ok := used[i]; ok {
			continue
		}
		used[i] = true
		*path = append(*path, tiles[i])
		backtrack(number, tiles, path, filter, used, res)
		*path = (*path)[:len(*path)-1]
		used[i] = false
	}
}