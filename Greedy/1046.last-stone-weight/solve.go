import "sort"

func lastStoneWeight(stones []int) int {
	for len(stones) > 1 {
		sort.Ints(stones)
		big := stones[len(stones)-1]
		second := stones[len(stones)-2]
		switch {
		case big == second:
			stones = stones[:len(stones)-2]
		default:
			stones[len(stones)-2] = big - second
			stones = stones[:len(stones)-1]
		}
	}
	if len(stones) == 1 {
		return stones[0]
	}
	return 0
}