func twoSum(nums []int, target int) []int {
	helper := map[int]int{}
	res := []int{}
	for i := 0; i < len(nums); i++ {
		c := target - nums[i]
		if idx, ok := helper[c]; ok {
			res = append(res, idx, i)
			return res
		}
		helper[nums[i]] = i
	}
	return res
}