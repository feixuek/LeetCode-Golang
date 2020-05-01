import "errors"

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

func (s *Stack) Top() (elem int, err error) {
	if s.Empty() {
		err = errors.New("empty Stack")
		return
	}
	elem = s.Data[s.Len-1]
	return
}

func (s *Stack) Empty() bool {
	return s.Len == 0
}

type MyQueue struct {
	Data   *Stack
	Helper *Stack
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	mq := MyQueue{
		Data:   &Stack{},
		Helper: &Stack{},
	}
	return mq
}

/** Push element x to the back of queue. */
func (mq *MyQueue) Push(elem int) {
	mq.Data.Push(elem)
}

/** Removes the element from in front of queue and returns that element. */
func (mq *MyQueue) Pop() int {
	if !mq.Helper.Empty() {
		elem, _ := mq.Helper.Pop()
		return elem
	}
	for !mq.Data.Empty() {
		elem, _ := mq.Data.Pop()
		mq.Helper.Push(elem)
	}
	elem, _ := mq.Helper.Pop()
	return elem
}

/** Get the front element. */
func (mq *MyQueue) Peek() int {
	if !mq.Helper.Empty() {
		elem, _ := mq.Helper.Top()
		return elem
	}
	for !mq.Data.Empty() {
		elem, _ := mq.Data.Pop()
		mq.Helper.Push(elem)
	}
	elem, _ := mq.Helper.Top()
	return elem
}

/** Returns whether the queue is empty. */
func (mq *MyQueue) Empty() bool {
	return mq.Data.Empty() && mq.Helper.Empty()
}