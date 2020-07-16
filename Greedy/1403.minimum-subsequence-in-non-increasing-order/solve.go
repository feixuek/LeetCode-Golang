import "sort"

func minSubsequence(nums []int) []int {
	sort.Ints(nums)
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	sum /= 2
	res := []int{}
	for i := len(nums) - 1; i >= 0; i-- {
		res = append(res, nums[i])
		sum -= nums[i]
		if sum < 0 {
			break
		}
	}
	return res
}