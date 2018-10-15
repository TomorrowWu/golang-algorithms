package stack

// Stack defines a stack interface
type Stack interface {
	Push(v interface{})
	Pop() interface{}
	IsEmpty() bool
	Top() interface{}
	Flush()
}
