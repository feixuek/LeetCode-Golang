func permute(nums []int) [][]int {
	res := make([][]int, 0)
	filter := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		filter[nums[i]] = false
	}
	backtrace(nums, &[]int{}, &res, filter)
	return res
}

func backtrace(nums []int, trace *[]int, res *[][]int, filter map[int]bool) {
	if len(nums) == len(*trace) {
		dst := make([]int, len(*trace))
		copy(dst, *trace)
		*res = append(*res, dst)
		return
	}
	for i := 0; i < len(nums); i++ {
		if exist := filter[nums[i]]; exist {
			continue
		}
		filter[nums[i]] = true
		*trace = append(*trace, nums[i])
		backtrace(nums, trace, res, filter)
		*trace = (*trace)[:len(*trace)-1]
		filter[nums[i]] = false
	}
}