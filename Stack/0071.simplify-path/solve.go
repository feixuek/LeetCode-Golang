import "strings"

type Stack struct {
	Data []string
}

func (s *Stack) Push(elem string) {
	s.Data = append(s.Data, elem)
}

func (s *Stack) Pop() (elem string) {
	if s.Empty() {
		return
	}
	elem = s.Data[s.Len()-1]
	s.Data = s.Data[:s.Len()-1]
	return
}

func (s *Stack) Top() (elem string) {
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

func simplifyPath(path string) string {
	stack := new(Stack)
	set := map[string]bool{
		".":  true,
		"..": true,
		"":   true,
	}
	arr := strings.Split(path, "/")
	for _, v := range arr {
		switch {
		case v == ".." && !stack.Empty():
			stack.Pop()
		case !set[v]:
			stack.Push(v)
		}
	}
	res := make([]string, stack.Len())
	for i := stack.Len() - 1; i >= 0; i-- {
		res[i] = stack.Pop()
	}
	return "/" + strings.Join(res, "/")
}