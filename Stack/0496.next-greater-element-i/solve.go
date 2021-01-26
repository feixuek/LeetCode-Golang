func nextGreaterElement(nums1 []int, nums2 []int) []int {
	filter := make(map[int]int)
	stack := []int{}
	for i := len(nums2) - 1; i >= 0; i-- {
		for len(stack) > 0 && stack[len(stack)-1] <= nums2[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			filter[nums2[i]] = -1
		} else {
			filter[nums2[i]] = stack[len(stack)-1]
		}
		stack = append(stack, nums2[i])
	}
	res := []int{}
	for i := 0; i < len(nums1); i++ {
		res = append(res, filter[nums1[i]])
	}
	return res
}