type NumArray struct {
	elem [][]int
}

func Constructor(nums []int) NumArray {
	var elem [][]int
	for i := 0; i < len(nums); i++ {
		tmp := make([]int, len(nums))
		elem = append(elem, tmp)
	}
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			if j-1 < 0 {
				elem[i][j] = nums[i]
				continue
			}
			elem[i][j] = elem[i][j-1] + nums[j]
		}
	}
	ret := NumArray{elem: elem}
	for i := 0; i < len(elem); i++ {
		for j := 0; j < len(elem[i]); j++ {
		}
	}
	return ret
}

func (this *NumArray) SumRange(i int, j int) int {
	return this.elem[i][j]
}