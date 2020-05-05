func repeatedNTimes(A []int) int {
	for i := 0; i < len(A)-2; i++ {
		if A[i] == A[i+1] || A[i] == A[i+2] {
			return A[i]
		}
	}
	return A[len(A)-1]
}