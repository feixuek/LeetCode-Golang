type Stack struct {
	Data []int
}

func (s *Stack) Push(elem int) {
	s.Data = append(s.Data, elem)
}

func (s *Stack) Pop() (elem int) {
	if s.Empty() {
		return
	}
	elem = s.Data[s.Len()-1]
	s.Data = s.Data[:s.Len()-1]
	return
}

func (s *Stack) Top() (elem int) {
	if s.Empty() {
		return
	}
	elem = s.Data[s.Len()-1]
	return
}

func (s *Stack) Empty() bool {
	return len(s.Data) == 0
}

func (s *Stack) Len() int {
	return len(s.Data)
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	stack := new(Stack)
	set := make(map[int]int)
	for i := len(nums2) - 1; i >= 0; i-- {
		for !stack.Empty() && nums2[i] > stack.Top() {
			stack.Pop()
		}
		if stack.Empty() {
			set[nums2[i]] = -1
		} else {
			set[nums2[i]] = stack.Top()
		}
		stack.Push(nums2[i])
	}
	ret := make([]int, len(nums1))
	for i := 0; i < len(nums1); i++ {
		if v, ok := set[nums1[i]]; ok {
			ret[i] = v
		} else {
			ret[i] = -1
		}
	}
	return ret
}