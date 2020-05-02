import (
	"errors"
	"strconv"
)

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

func calPoints(ops []string) int {
	stack := new(Stack)
	for _, v := range ops {
		switch v {
		case "+":
			elem1, _ := stack.Pop()
			elem := stack.Top() + elem1
			stack.Push(elem1)
			stack.Push(elem)
		case "C":
			stack.Pop()
		case "D":
			stack.Push(stack.Top() * 2)
		default:
			elem, _ := strconv.Atoi(v)
			stack.Push(elem)
		}
	}
	sum := 0
	for !stack.Empty() {
		elem, _ := stack.Pop()
		sum += elem
	}
	return sum
}