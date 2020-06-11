func fib(N int) int {
	if N <= 0 {
		return 0
	}
	if N == 1 || N == 2 {
		return 1
	}
	pre, curr := 1, 1
	for i := 3; i <= N; i++ {
		sum := pre + curr
		pre = curr
		curr = sum
	}
	return curr
}