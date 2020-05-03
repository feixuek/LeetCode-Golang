func singleNumber(nums []int) int {
	var ret int
	for _, v := range nums {
		ret ^= v
	}
	return ret
}