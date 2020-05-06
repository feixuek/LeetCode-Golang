func findTheDifference(s string, t string) byte {
	helper := make(map[rune]int)
	for _, v := range s {
		helper[v]++
	}
	var ret rune
	for _, v := range t {
		count, ok := helper[v]
		if !ok || count <= 0 {
			ret = v
			break
		}
		helper[v]--
	}
	return byte(ret)
}