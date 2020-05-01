import "errors"

type Stack struct {
	Data []int
	Len  int64
}

func (s *Stack) Push(elem int) {
	s.Data = append(s.Data, elem)
	s.Len++
}

func (s *Stack) Pop() (elem int, err error) {
	if s.Empty() {
		err = errors.New("empty Stack")
		return
	}
	elem = s.Data[s.Len-1]
	s.Data = s.Data[:s.Len-1]
	s.Len--
	return
}

func (s *Stack) Top() (elem int) {
	if s.Empty() {
		elem = -1
		return
	}
	elem = s.Data[s.Len-1]
	return
}

func (s *Stack) Empty() bool {
	return s.Len == 0
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	stack := new(Stack)
	helper := make(map[int]int)
	for i := len(nums2) - 1; i >= 0; i-- {
		for !stack.Empty() && stack.Top() <= nums2[i] {
			stack.Pop()
		}
		helper[nums2[i]] = stack.Top()
		stack.Push(nums2[i])
	}
	ret := []int{}
	for i := 0; i < len(nums1); i++ {
		if v, ok := helper[nums1[i]]; ok {
			ret = append(ret, v)
		} else {
			ret = append(ret, -1)
		}
	}
	return ret
}