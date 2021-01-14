import (
	"fmt"
	"sort"
)

func subsetsWithDup(nums []int) [][]int {

	res := make([][]int, 0)
	filter := make(map[int]bool)
	filter2 := make(map[string]bool)
	backtrace(nums, &res, &[]int{}, &filter, 0, &filter2)
	return res
}

func backtrace(nums []int, res *[][]int, trace *[]int, filter *map[int]bool, start int, filter2 *map[string]bool) {
	dst := make([]int, len(*trace))
	copy(dst, *trace)
	sort.Ints(dst)
	cache := fmt.Sprintf("%v", dst)
	if ok := (*filter2)[cache]; !ok {
		*res = append(*res, dst)
		(*filter2)[cache] = true
	}
	for i := start; i < len(nums); i++ {
		*trace = append(*trace, nums[i])
		backtrace(nums, res, trace, filter, i+1, filter2)
		*trace = (*trace)[:len(*trace)-1]
	}
}