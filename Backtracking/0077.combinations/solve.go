func combine(n int, k int) [][]int {
	if n <= 0 || k <= 0 {
		return nil
	}
	result := [][]int{}
	trace := []int{}
	var f func(start int)
	f = func(start int) {
		if len(trace) == k {
			v := make([]int, k)
			copy(v, trace)
			result = append(result, v)
		}
		for i := start; i <= n; i++ {
			trace = append(trace, i)
			f(i + 1)
			trace = trace[:len(trace)-1]
		}
	}
	f(1)
	return result
}