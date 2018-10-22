package src

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

// MyQueue defines Stack1 queue by two stacks
type MyQueue struct {
	Stack1, Stack2 *stack
}

// Constructor Initialize your data structure here.
func Constructor() MyQueue {
	return MyQueue{
		Stack1: newStack(),
		Stack2: newStack(),
	}
}

// Push element x to the back of queue.
func (queue *MyQueue) Push(x int) {
	queue.Stack1.push(x)
}

// Pop Removes the element from in front of queue and returns that element.
func (queue *MyQueue) Pop() int {
	if queue.Stack2.isEmpty() {
		//优化: 栈a中留一个元素供pop,可以少一次操作
		for queue.Stack1.len() > 1 {
			queue.Stack2.push(queue.Stack1.pop())
		}
		return queue.Stack1.pop()
	}
	return queue.Stack2.pop()
}

// Peek Get the front element.
func (queue *MyQueue) Peek() int {
	res := queue.Pop()
	//队列为空
	if res != -1 {
		queue.Stack2.push(res)
	}
	return res
}

// Empty Returns whether the queue is empty.
func (queue *MyQueue) Empty() bool {
	return queue.Stack1.isEmpty() && queue.Stack2.isEmpty()
}

// stack defines Stack1 stack
type stack struct {
	nums []int
}

// newStack creates a empty stack
func newStack() *stack {
	return &stack{
		nums: []int{},
	}
}

func (s *stack) push(n int) {
	s.nums = append(s.nums, n)
}

func (s *stack) pop() int {
	if s.isEmpty() {
		return -1
	}
	res := s.nums[len(s.nums)-1]
	s.nums = s.nums[:len(s.nums)-1]
	return res
}

func (s *stack) len() int {
	return len(s.nums)
}
func (s *stack) isEmpty() bool {
	return s.len() == 0
}
