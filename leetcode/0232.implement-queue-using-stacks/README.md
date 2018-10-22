### [232. 用栈实现队列](https://leetcode-cn.com/problems/implement-queue-using-stacks/description/)
使用栈实现队列的下列操作：
- push(x) -- 将一个元素放入队列的尾部。
- pop() -- 从队列首部移除元素。
- peek() -- 返回队列首部的元素。
- empty() -- 返回队列是否为空。

**示例:**
```
MyQueue queue = new MyQueue();

queue.push(1);
queue.push(2);
queue.peek();  // 返回 1
queue.pop();   // 返回 1
queue.empty(); // 返回 false
```

**说明:**
- 你只能使用标准的栈操作 -- 也就是只有 push to top, peek/pop from top, size, 和 is empty 操作是合法的。
- 你所使用的语言也许不支持栈。你可以使用 list 或者 deque（双端队列）来模拟一个栈，只要是标准的栈操作即可。
- 假设所有操作都是有效的 （例如，一个空的队列不会调用 pop 或者 peek 操作）。

### 解题思路
- 在Golang中,本身未提供栈数据结构,先基于数组实现栈结构
- 优化: 入队时,将元素压入s1,出队时，判断s2是否为空，如不为空，则直接弹出顶元素；如为空，则将s1的元素逐个“倒入”s2，把最后一个元素弹出并出队
- 避免反复"倒"栈,仅在需要时才"倒"一次
- 细节: 考虑没有元素可供出队时的处理（2个栈都为空的时候，出队操作一定会引起异常）

### 代码实现

```Golang
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
	if queue.Stack2.isEmpty() {
		if queue.Stack1.isEmpty() {
			return -1
		}
		return queue.Stack1.nums[0]
	}
	return queue.Stack2.nums[queue.Stack2.len()-1]
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
```

### GitHub
- [源码在这里](https://github.com/TomorrowWu/golang-algorithms/blob/master/leetcode/0232.implement-queue-using-stacks/src/queue-by-stacks.go)
- 项目中会提供各种数据结构及算法的Golang实现,以及 LeetCode 解法

##### 参考资料

[用两个栈实现一个队列——我作为面试官的小结](https://www.cnblogs.com/wanghui9072229/archive/2011/11/22/2259391.html)
