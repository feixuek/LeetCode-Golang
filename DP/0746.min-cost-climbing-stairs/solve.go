/*
func minCostClimbingStairs(cost []int) int {
	length := len(cost)
	switch {
	case length == 0:
		return 0
	case length == 1:
		return cost[0]
	default:
		res := make([]int, length)
		res[0] = cost[0]
		res[1] = cost[1]
		for i := 2; i < length; i++ {
			res[i] = min(res[i-1], res[i-2]) + cost[i]
		}
		return min(res[length-1], res[length-2])
	}
}
*/
func minCostClimbingStairs(cost []int) int {
	length := len(cost)
	switch {
	case length == 0:
		return 0
	case length == 1:
		return cost[0]
	default:
		res0 := cost[0]
		res1 := cost[1]
		for i := 2; i < length; i++ {
			tmp := res1
			res1 = min(res1, res0) + cost[i]
			res0 = tmp
		}
		return min(res0, res1)
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}