func balancedStringSplit(s string) int {
	res, r, l := 0, 0, 0
	for _, v := range s {
		switch v {
		case 'R':
			if r == 0 && l == 0 {
				res++
			}
			if l > 0 {
				l--
			} else {
				r++
			}
		case 'L':
			if r == 0 && l == 0 {
				res++
			}
			if r > 0 {
				r--
			} else {
				l++
			}
		}
	}
	return res
}