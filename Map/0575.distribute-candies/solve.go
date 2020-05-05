func distributeCandies(candies []int) int {
	helper := make(map[int]int)
	for i := 0; i < len(candies); i++ {
		helper[candies[i]]++
	}
	length := len(candies) / 2
	category := 0
	for _, v := range helper {
		if v > 0 {
			category++
			length--
			if length == 0 {
				break
			}
		}
	}
	return category
}
