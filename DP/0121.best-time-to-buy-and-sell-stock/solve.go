/*
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	var res [][]int
	for i := 0; i < len(prices); i++ {
		tmp := make([]int, 2)
		res = append(res, tmp)
	}
	for i := 0; i < len(prices); i++ {
		if i-1 < 0 {
			res[i][0] = 0
			res[i][1] = -prices[i]
			continue
		}
		res[i][0] = max(res[i-1][0], res[i-1][1]+prices[i])
		res[i][1] = max(res[i-1][1], -prices[i])
	}
	return res[len(prices)-1][0]
}
*/
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	res_i_0, res_i_1 := 0, 0
	for i := 0; i < len(prices); i++ {
		if i-1 < 0 {
			res_i_0 = 0
			res_i_1 = -prices[i]
			continue
		}
		res_i_0 = max(res_i_0, res_i_1+prices[i])
		res_i_1 = max(res_i_1, -prices[i])
	}
	return res_i_0
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}