import (
	"fmt"
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	filter := make(map[string]bool)
	backtrace(candidates, target, &res, &[]int{}, &filter)
	return res
}

func backtrace(candidates []int, target int, res *[][]int, trace *[]int, filter *map[string]bool) {
	if target < 0 {
		return
	}
	if target == 0 {
		dst := make([]int, len(*trace))
		copy(dst, *trace)
		sort.Ints(dst)
		cache := fmt.Sprintf("%v", dst)
		if ok := (*filter)[cache]; ok {
			return
		}
		(*filter)[cache] = true
		*res = append(*res, dst)
		return
	}
	for i := 0; i < len(candidates); i++ {
		*trace = append(*trace, candidates[i])
		backtrace(candidates, target-candidates[i], res, trace, filter)
		*trace = (*trace)[:len(*trace)-1]
	}
}