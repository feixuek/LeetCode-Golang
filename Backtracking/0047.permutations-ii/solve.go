import "strconv"

func permuteUnique(nums []int) [][]int {
	res := make([][]int, 0)
	filter := make(map[int]bool)
	filter2 := make(map[string]bool)
	for i := 0; i < len(nums); i++ {
		filter[i] = false
	}
	backtrace(nums, &res, &[]int{}, filter, &filter2)
	return res
}

func backtrace(nums []int, res *[][]int, trace *[]int, filter map[int]bool, filter2 *map[string]bool) {
	if len(nums) == len(*trace) {
		dst := make([]int, len(nums))
		copy(dst, *trace)
		cache := ""
		for i := 0; i < len(dst); i++ {
			cache += strconv.Itoa(dst[i])
		}
		if (*filter2)[cache] {
			return
		}
		(*filter2)[cache] = true
		*res = append(*res, dst)
		return
	}
	for i := 0; i < len(nums); i++ {
		if exist, ok := filter[i]; ok && exist {
			continue
		}
		filter[i] = true
		*trace = append(*trace, nums[i])
		backtrace(nums, res, trace, filter, filter2)
		filter[i] = false
		*trace = (*trace)[:len(*trace)-1]
	}
}