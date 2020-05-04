func jump(nums []int) int {
	steps, currEnd, currFarthest := 0, 0, 0
	for i := 0; i < len(nums)-1; i++ {
		currFarthest = max(currFarthest, i+nums[i])
		if i == currEnd {
			steps++
			currEnd = currFarthest
		}
	}
	return steps
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}