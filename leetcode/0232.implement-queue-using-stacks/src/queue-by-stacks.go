package src

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

type MyQueue struct {
}

/** Initialize your data structure here. */
func Constructor() MyQueue {

	return MyQueue{}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {

}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {

	return -1
}

/** Get the front element. */
func (this *MyQueue) Peek() int {

	return -1
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {

	return false
}
