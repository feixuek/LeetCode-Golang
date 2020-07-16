func maxProfit(prices []int) int {
	res := 0
	for i := 1; i < len(prices); i++ {
		v := prices[i] - prices[i-1]
		if v > 0 {
			res += v
		}
	}
	return res
}