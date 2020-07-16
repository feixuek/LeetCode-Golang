import "sort"

func twoCitySchedCost(costs [][]int) int {
	sum := 0
	sort.Slice(costs, func(i, j int) bool { return costs[i][0]-costs[i][1] < costs[j][0]-costs[j][1] })
	length := len(costs) / 2
	for i := 0; i < length; i++ {
		sum += costs[i][0] + costs[i+length][1]
	}
	return sum
}