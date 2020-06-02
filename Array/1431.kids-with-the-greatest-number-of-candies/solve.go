func kidsWithCandies(candies []int, extraCandies int) []bool {
	max := candies[0]
	for i := 1; i < len(candies); i++ {
		if candies[i] > max {
			max = candies[i]
		}
	}
	ret := make([]bool, len(candies))
	for i := 0; i < len(candies); i++ {
		if candies[i]+extraCandies >= max {
			ret[i] = true
		}
	}
	return ret
}