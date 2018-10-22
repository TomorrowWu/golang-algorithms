package src

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

// MyQueue defines a queue by two stacks
type MyQueue struct {
	a, b *stack
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		a: newStack(),
		b: newStack(),
	}
}

// Push element x to the back of queue.
func (queue *MyQueue) Push(x int) {
	queue.a.push(x)
}

// Pop Removes the element from in front of queue and returns that element.
func (queue *MyQueue) Pop() int {
	if queue.b.isEmpty() {
		//优化: 栈a中留一个元素供pop,可以少一次操作
		for queue.a.len() > 1 {
			queue.b.push(queue.a.pop())
		}
		return queue.a.pop()
	}
	return queue.b.pop()
}

// Peek Get the front element.
func (queue *MyQueue) Peek() int {
	if queue.b.isEmpty() {
		if queue.a.isEmpty() {
			return -1
		} else {
			return queue.a.nums[0]
		}
	}
	return queue.b.nums[queue.b.len()-1]
}

// Empty Returns whether the queue is empty.
func (queue *MyQueue) Empty() bool {
	return queue.a.isEmpty() && queue.b.isEmpty()
}

// stack defines a stack
type stack struct {
	nums []int
}

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
