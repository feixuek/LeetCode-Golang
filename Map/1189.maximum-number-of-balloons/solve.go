func maxNumberOfBalloons(text string) int {
	helper := make(map[rune]int)
	for _, v := range text {
		switch v {
		case 'b', 'a', 'l', 'o', 'n':
			helper[v]++
		}
	}
	ret := 1<<32 - 1
	for _, v := range "balon" {
		n, ok := helper[v]
		if !ok {
			return 0
		}
		switch v {
		case 'l', 'o':
			n /= 2
		}
		if n < ret {
			ret = n
		}
	}
	return ret
}