package stack

import "fmt"

type node struct {
	next *node
	val  interface{}
}

// LinkedListStack defines a stack based on linked-list
type LinkedListStack struct {
	topNode *node //top of stack
}

// NewLinkedListStack creates a empty stack
func NewLinkedListStack() *LinkedListStack {
	return &LinkedListStack{
		topNode: nil,
	}
}

// IsEmpty returns whether stack is empty
func (s *LinkedListStack) IsEmpty() bool {
	return s.topNode == nil
}

// Push pushes a element to stack top
func (s *LinkedListStack) Push(v interface{}) {
	s.topNode = &node{
		next: s.topNode,
		val:  v,
	}
}

// Pop pops the top element of stack
func (s *LinkedListStack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	val := s.topNode.val
	s.topNode = s.topNode.next
	return val
}

// Top returns the top element of stack
func (s *LinkedListStack) Top() interface{} {
	if s.IsEmpty() {
		return nil
	}
	return s.topNode.val
}

// Flush flushes the stack
func (s *LinkedListStack) Flush() {
	s.topNode = nil
}

// Print prints the stack
func (s *LinkedListStack) Print() {
	if s.IsEmpty() {
		fmt.Println("empty stack")
	} else {
		cur := s.topNode
		for nil != cur {
			fmt.Println(cur.val)
			cur = cur.next
		}
	}
}
