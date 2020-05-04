func countCharacters(words []string, chars string) int {
	helper := make(map[rune]int)
	ret := 0
	for _, v := range chars {
		helper[v]++
	}
	for i := 0; i < len(words); i++ {
		tmp := make(map[rune]int)
		for _, v := range words[i] {
			tmp[v]++
		}
		for v, n := range tmp {
			if hn, ok := helper[v]; !ok || hn < n {
				goto LOOP
			}
		}
		ret += len(words[i])
	LOOP:
	}
	return ret
}