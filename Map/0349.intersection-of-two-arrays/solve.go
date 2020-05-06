func intersection(nums1 []int, nums2 []int) []int {
	helper := make(map[int]int)
	for _, v := range nums1 {
		helper[v]++
	}
	res := []int{}
	for _, v := range nums2 {
		if _, ok := helper[v]; ok {
			res = append(res, v)
			delete(helper, v)
		}
	}
	return res
}