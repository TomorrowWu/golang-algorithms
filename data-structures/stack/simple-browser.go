package stack

import "fmt"

// Browser defines a container to store browsing history
type Browser struct {
	forwardStack Stack
	backStack    Stack
}

// NewBrowser creates a browser
func NewBrowser() *Browser {
	return &Browser{
		forwardStack: NewArrayStack(),
		backStack:    NewLinkedListStack(),
		//backStack: NewArrayStack(),
	}
}

// CanForward returns if browser can forward
func (b *Browser) CanForward() bool {
	if b.forwardStack.IsEmpty() {
		return false
	}
	return true
}

// CanBack returns if browser can back
func (b *Browser) CanBack() bool {
	if b.backStack.IsEmpty() {
		return false
	}
	return true
}

// Open opens an addr
func (b *Browser) Open(addr string) {
	fmt.Printf("Open new addr %+v\n", addr)
	b.forwardStack.Flush()
}

// PushBackStack pushes a addr into backStack
func (b *Browser) PushBackStack(addr string) {
	b.backStack.Push(addr)
}

// Forward forwards to addr
func (b *Browser) Forward() {
	if b.forwardStack.IsEmpty() {
		return
	}
	top := b.forwardStack.Pop()
	b.backStack.Push(top)
	fmt.Printf("forward to %+v\n", top)
}

// Back backs to addr
func (b *Browser) Back() {
	if b.backStack.IsEmpty() {
		return
	}

	top := b.backStack.Pop()
	b.forwardStack.Push(top)
	fmt.Printf("back to %+v\n", top)
}
