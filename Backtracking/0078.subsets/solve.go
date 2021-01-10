func subsets(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	result := [][]int{}
	trace := []int{}
	var f func(start int)
	f = func(start int) {
		v := make([]int, len(trace))
		copy(v, trace)
		result = append(result, v)
		for i := start; i < len(nums); i++ {
			trace = append(trace, nums[i])
			f(i + 1)
			trace = trace[:len(trace)-1]
		}
	}
	f(0)
	return result
}