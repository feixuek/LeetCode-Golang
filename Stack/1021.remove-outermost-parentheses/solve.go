import "errors"

type Stack struct {
	Data []byte
	Len  int64
}

func (s *Stack) Push(elem byte) {
	s.Data = append(s.Data, elem)
	s.Len++
}

func (s *Stack) Pop() (elem byte, err error) {
	if s.Empty() {
		err = errors.New("empty Stack")
		return
	}
	elem = s.Data[s.Len-1]
	s.Data = s.Data[:s.Len-1]
	s.Len--
	return
}

func (s *Stack) Top() (elem byte) {
	if s.Empty() {
		return
	}
	elem = s.Data[s.Len-1]
	return
}

func (s *Stack) Empty() bool {
	return s.Len == 0
}

func removeOuterParentheses(S string) string {
	stack := new(Stack)
	start, end := 0, 0
	ret := ""
	for i := 0; i < len(S); i++ {
		switch S[i] {
		case '(':
			if stack.Empty() {
				start = i
			}
			stack.Push(S[i])
		case ')':
			stack.Pop()
			if stack.Empty() {
				end = i
			}
		}
		if stack.Empty() {
			ret += string(S[start+1 : end])
		}
	}
	return ret
}