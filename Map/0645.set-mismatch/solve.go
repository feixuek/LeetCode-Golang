func findErrorNums(nums []int) []int {
	res := make([]int, 2)
	helper := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		helper[i+1] = 1
	}
	for i := 0; i < len(nums); i++ {
		if _, ok := helper[nums[i]]; ok {
			helper[nums[i]]--
		}
	}
	for k, v := range helper {
		switch v {
		case 1:
			res[1] = k
		case -1:
			res[0] = k
		}
	}
	return res
}