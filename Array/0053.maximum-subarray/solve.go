func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sum, curr := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		curr = max(curr+nums[i], nums[i])
		sum = max(curr, sum)
	}
	return sum
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}