func intersect(nums1 []int, nums2 []int) []int {
	helper := make(map[int]int)
	for _, v := range nums1 {
		helper[v]++
	}
	res := []int{}
	for _, v := range nums2 {
		if count, ok := helper[v]; ok {
			if count > 0 {
				res = append(res, v)
				helper[v]--
			}
		}
	}
	return res
}