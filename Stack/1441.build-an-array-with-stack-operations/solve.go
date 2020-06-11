func buildArray(target []int, n int) []string {
	idx := 0
	res := []string{}
	for i := 1; i <= n; i++ {
		if idx >= len(target) {
			break
		}
		switch {
		case target[idx] == i:
			res = append(res, "Push")
			idx++
		default:
			res = append(res, []string{"Push", "Pop"}...)
		}
	}
	return res
}