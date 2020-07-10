import "strconv"

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

func evalRPN(tokens []string) int {
	stack := new(Stack)
	for _, v := range tokens {
		switch v {
		case "+":
			v1, v2 := stack.Pop(), stack.Pop()
			stack.Push(v1 + v2)
		case "-":
			v1, v2 := stack.Pop(), stack.Pop()
			stack.Push(v2 - v1)
		case "*":
			v1, v2 := stack.Pop(), stack.Pop()
			stack.Push(v1 * v2)
		case "/":
			v1, v2 := stack.Pop(), stack.Pop()
			stack.Push(v2 / v1)
		default:
			if val, err := strconv.Atoi(v); err == nil {
				stack.Push(val)
			}
		}
	}
	return stack.Pop()
}