func findDuplicate(nums []int) int {
	for len(nums) > 0 {
		slow, fast := nums[0], nums[nums[0]]
		for slow != fast {
			slow = nums[slow]
			fast = nums[nums[fast]]
		}
		slow = 0
		for slow != fast {
			slow = nums[slow]
			fast = nums[fast]
		}
		return slow
	}
	return -1
}