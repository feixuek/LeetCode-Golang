import "sort"

func largestSumAfterKNegations(A []int, K int) int {
	for i := 0; i < K; i++ {
		sort.Ints(A)
		A[0] *= -1
	}
	sum := 0
	for i := 0; i < len(A); i++ {
		sum += A[i]
	}
	return sum
}