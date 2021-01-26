func removeDuplicates(nums []int) int {
	left, right := 0, 0
	for right < len(nums) {
		if nums[left] != nums[right] {
			left++
			nums[left] = nums[right]
		}
		right++
	}
	return left + 1
}