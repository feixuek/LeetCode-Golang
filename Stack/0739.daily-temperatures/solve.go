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

func dailyTemperatures(T []int) []int {
	stack := new(Stack)
	res := make([]int, len(T))
	for i := len(T) - 1; i >= 0; i-- {
		for !stack.Empty() && T[i] >= T[stack.Top()] {
			stack.Pop()
		}
		if stack.Empty() {
			res[i] = 0
		} else {
			res[i] = stack.Top() - i
		}
		stack.Push(i)
	}
	return res
}