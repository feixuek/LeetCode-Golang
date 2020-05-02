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

func removeDuplicates(S string) string {
	stack := new(Stack)
	for _, v := range S {
		if stack.Top() == v {
			stack.Pop()
		} else {
			stack.Push(v)
		}
	}
	length := stack.Len
	runes := make([]rune, length)
	length--
	for !stack.Empty() {
		v, _ := stack.Pop()
		runes[length] = v
		length--
	}
	return string(runes)
}