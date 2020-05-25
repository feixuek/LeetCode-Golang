func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	res := make([]int, len(nums))
	res[0] = nums[0]
	res[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		res[i] = max(res[i-1], nums[i]+res[i-2])
	}
	return res[len(res)-1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}