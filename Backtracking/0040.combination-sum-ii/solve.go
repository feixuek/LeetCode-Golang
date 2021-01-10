import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	filter := make(map[string]bool)
	backtrace(candidates, target, 0, &res, &[]int{}, &filter)
	return res
}

func backtrace(candidates []int, target, start int, res *[][]int, trace *[]int, filter *map[string]bool) {
	if target < 0 {
		return
	}
	if target == 0 {
		dst := make([]int, len(*trace))
		copy(dst, *trace)
		sort.Ints(dst)
		cache := fmt.Sprintf("%v", dst)
		if (*filter)[cache] {
			return
		}
		(*filter)[cache] = true
		*res = append(*res, dst)
		return
	}
	for i := start; i < len(candidates); i++ {
		*trace = append(*trace, candidates[i])
		backtrace(candidates, target-candidates[i], i+1, res, trace, filter)
		*trace = (*trace)[:len(*trace)-1]
	}
}