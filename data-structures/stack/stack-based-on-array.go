package stack

import "fmt"

// ArrayStack defines the stack based on array
type ArrayStack struct {
	data []interface{} //data
	top  int           //The stack top index in data
}

// NewArrayStack creates a empty stack
func NewArrayStack() *ArrayStack {
	return &ArrayStack{
		data: make([]interface{}, 0, 32),
		top:  -1,
	}
}

// IsEmpty returns whether stack is empty
func (s *ArrayStack) IsEmpty() bool {
	return s.top < 0
}

// Push pushes a element to stack top
func (s *ArrayStack) Push(v interface{}) {
	if s.top < 0 {
		s.top = 0
	} else {
		s.top++
	}

	if s.top > len(s.data)-1 {
		s.data = append(s.data, v)
	} else {
		s.data[s.top] = v
	}
}

// Pop pops the top element of stack
func (s *ArrayStack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	topEle := s.data[s.top]
	s.top--
	return topEle
}

// Top returns the top element of stack
func (s *ArrayStack) Top() interface{} {
	if s.IsEmpty() {
		return nil
	}
	return s.data[s.top]
}

// Flush flushes the stack
func (s *ArrayStack) Flush() {
	s.top = -1
}

// Print prints the stack
func (s *ArrayStack) Print() {
	if s.IsEmpty() {
		fmt.Println("empty stack")
	} else {
		for i := s.top; i >= 0; i-- {
			fmt.Println(s.data[i])
		}
	}
}
