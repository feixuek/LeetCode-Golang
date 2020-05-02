import "errors"

type Stack struct {
	Data []rune
	Len  int64
}

func (s *Stack) Push(elem rune) {
	s.Data = append(s.Data, elem)
	s.Len++
}

func (s *Stack) Pop() (elem rune, err error) {
	if s.Empty() {
		err = errors.New("empty Stack")
		return
	}
	elem = s.Data[s.Len-1]
	s.Data = s.Data[:s.Len-1]
	s.Len--
	return
}

func (s *Stack) Top() (elem rune) {
	if s.Empty() {
		return
	}
	elem = s.Data[s.Len-1]
	return
}

func (s *Stack) Empty() bool {
	return s.Len == 0
}

func backspaceCompare(S string, T string) bool {
	stack1 := new(Stack)
	stack2 := new(Stack)
	for _, v := range S {
		switch v {
		case '#':
			stack1.Pop()
		default:
			stack1.Push(v)
		}
	}
	for _, v := range T {
		switch v {
		case '#':
			stack2.Pop()
		default:
			stack2.Push(v)
		}
	}
	if (stack1.Empty() && !stack2.Empty()) || (!stack1.Empty() && stack2.Empty()) {
		return false
	}
	for !stack1.Empty() {
		v1, _ := stack1.Pop()
		v2, err := stack2.Pop()
		if err != nil || v1 != v2 {
			return false
		}
	}
	if !stack2.Empty() {
		return false
	}
	return true
}