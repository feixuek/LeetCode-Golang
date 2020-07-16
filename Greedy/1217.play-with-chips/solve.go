func minCostToMoveChips(chips []int) int {
	odd, even := 0, 0
	for _, v := range chips {
		if v%2 == 0 {
			even++
		} else {
			odd++
		}
	}
	if even > odd {
		return odd
	}
	return even
}