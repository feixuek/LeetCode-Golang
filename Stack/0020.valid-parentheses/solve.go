import "errors"

type Stack struct {
	Data []byte
	Len  int64
}

func (s *Stack) Push(elem byte) {
	s.Data = append(s.Data, elem)
	s.Len++
}

func (s *Stack) Pop() (byte, error) {
	if s.Empty() {
		return '0', errors.New("empty stack")
	}
	elem := s.Data[s.Len-1]
	s.Data = s.Data[:s.Len-1]
	s.Len--
	return elem, nil
}

func (s *Stack) Empty() bool {
	return s.Len == 0
}

func isValid(s string) bool {
	stack := new(Stack)
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(', '[', '{':
			stack.Push(s[i])
		case ')':
			if elem, err := stack.Pop(); err != nil || elem != '(' {
				return false
			}
		case ']':
			if elem, err := stack.Pop(); err != nil || elem != '[' {
				return false
			}
		case '}':
			if elem, err := stack.Pop(); err != nil || elem != '{' {
				return false
			}
		default:
			return false
		}
	}
	if stack.Empty() {
		return true
	}
	return false
}