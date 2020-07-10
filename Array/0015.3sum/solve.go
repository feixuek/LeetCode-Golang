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

func threeSum(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	var res [][]int
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		twoSumRes := twoSum(nums, i+1, nums[i]*-1)
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