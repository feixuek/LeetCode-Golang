const (
	c = 1 << 32
)

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	ret := make([]int, amount+1)
	ret[0] = 0
	for i := 1; i < len(ret); i++ {
		ret[i] = c
	}
	for i := 0; i < len(ret); i++ {
		for _, v := range coins {
			sub := i - v
			if sub < 0 {
				continue
			}
			ret[i] = min(ret[i], ret[sub]+1)
		}
	}
	if ret[amount] == c {
		return -1
	}
	return ret[amount]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}