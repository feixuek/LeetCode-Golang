import "sort"

func twoSum(nums []int, index, target int) [][]int {
	if len(nums) == 0 || index >= len(nums) {
		return nil
	}
	start := index
	end := len(nums) - 1
	var res [][]int
	for start < end {
		sum := nums[start] + nums[end]
		left := nums[start]
		right := nums[end]
		switch {
		case sum > target:
			for start < end && nums[end] == right {
				end--
			}
		case sum < target:
			for start < end && nums[start] == left {
				start++
			}
		default:
			res = append(res, []int{nums[start], nums[end]})
			for start < end && nums[start] == left {
				start++
			}
			for start < end && nums[end] == right {
				end--
			}
		}
	}
	return res
}

func threeSum(nums []int, index, target int) [][]int {
	if len(nums) == 0 || index >= len(nums) {
		return nil
	}
	var res [][]int
	for i := index; i < len(nums); i++ {
		twoSumRes := twoSum(nums, i+1, target-nums[i])
		for _, v := range twoSumRes {
			v = append(v, nums[i])
			res = append(res, v)
		}
		for i+1 < len(nums) && nums[i] == nums[i+1] {
			i++
		}
	}
	return res
}

func fourSum(nums []int, target int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	sort.Ints(nums)
	var res [][]int
	for i := 0; i < len(nums); i++ {
		threeSumRes := threeSum(nums, i+1, target-nums[i])
		for _, v := range threeSumRes {
			v = append(v, nums[i])
			res = append(res, v)
		}
		for i+1 < len(nums) && nums[i] == nums[i+1] {
			i++
		}
	}
	return res
}