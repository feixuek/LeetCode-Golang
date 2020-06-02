import "strconv"

type Stack struct {
	Data []byte
}

func (s *Stack) Push(elem byte) {
	s.Data = append(s.Data, elem)
}

func (s *Stack) Pop() (elem byte) {
	if s.Empty() {
		return
	}
	elem = s.Data[s.Len()-1]
	s.Data = s.Data[:s.Len()-1]
	return
}

func (s *Stack) Top() (elem byte) {
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

func decodeString(s string) string {
	stack := new(Stack)
	for i := 0; i < len(s); i++ {
		switch {
		case s[i] == ']':
			tmp := make([]byte, 0)
			for !stack.Empty() {
				elem := stack.Pop()
				switch {
				case elem == '[':
					number := string(stack.Pop())
					for stack.Top() >= '0' && stack.Top() <= '9' {
						number = string(stack.Pop()) + number
					}
					n, _ := strconv.Atoi(string(number))
					for i := 0; i < n; i++ {
						for j := len(tmp) - 1; j >= 0; j-- {
							stack.Push(tmp[j])
						}
					}
					goto LOOP
				default:
					tmp = append(tmp, elem)
				}
			}
		default:
			stack.Push(s[i])
		}
	LOOP:
	}
	res := make([]byte, stack.Len())
	for !stack.Empty() {
		elem := stack.Pop()
		res[stack.Len()] = elem
	}
	return string(res)
}