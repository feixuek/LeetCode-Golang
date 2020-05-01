type MinStack struct {
	Data []int
	Min  []int
	Len  int64
}

/** initialize your data structure here. */
func Constructor() MinStack {
	var ms MinStack
	return ms
}

func (ms *MinStack) Push(x int) {
	ms.Data = append(ms.Data, x)
	if ms.Empty() {
		ms.Min = append(ms.Min, x)
	} else {
		min := x
		if ms.GetMin() < x {
			min = ms.GetMin()
		}
		ms.Min = append(ms.Min, min)
	}
	ms.Len++
}

func (ms *MinStack) Pop() {
	ms.Data = ms.Data[:ms.Len-1]
	ms.Min = ms.Min[:ms.Len-1]
	ms.Len--
}

func (ms *MinStack) Top() int {
	return ms.Data[ms.Len-1]
}

func (ms *MinStack) GetMin() int {
	return ms.Min[ms.Len-1]
}

func (ms *MinStack) Empty() bool {
	return ms.Len == 0
}