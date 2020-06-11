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

func nextGreaterElements(nums []int) []int {
	stack := new(Stack)
	ret := make([]int, len(nums))
	for i := len(nums)*2 - 1; i >= 0; i-- {
		for !stack.Empty() && nums[i%len(nums)] >= stack.Top() {
			stack.Pop()
		}
		if stack.Empty() {
			ret[i%len(nums)] = -1
		} else {
			ret[i%len(nums)] = stack.Top()
		}
		stack.Push(nums[i%len(nums)])
	}
	return ret
}