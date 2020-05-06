func smallerNumbersThanCurrent2(nums []int) []int {
	var res []int
	helper := make(map[int]int)
	for _, v := range nums {
		helper[v]++
	}
	for _, v := range nums {
		n := 0
		for key, count := range helper {
			if key < v {
				n += count
			}
		}
		res = append(res, n)
	}
	return res
}