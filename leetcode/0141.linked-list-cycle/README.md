### 题目描述
给定一个链表，判断链表中是否有环。

**进阶:**
```
你能否不使用额外空间解决此题？
```

### 解题思路
- 无环链表，最后一个节点为nil，有环链表可以无限循环next下去
- 不用额外空间：快慢节点，慢节点一次走一步，快节点一次走两步，当进入环中，每次循环，快节点会离慢节点近一步，快节点最终会追上慢节点
- 用额外空间： 用map存走过的节点，第一个走过的节点就是环的入口

### 代码实现
```Golang
// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head != nil {
		slow := head
		fast := head
		for fast != nil && fast.Next != nil {
			slow = slow.Next
			fast = fast.Next.Next
			if slow == fast {
				return true
			}
		}
	}
	return false
}
```

### GitHub
- [源码传送门](https://github.com/TomorrowWu/golang-algorithms/blob/master/leetcode/0141.linked-list-cycle/src/linked-list-cycle.go)
- 项目中会提供各种数据结构及算法的Golang实现, LeetCode解题思路及答案

###### 题目来源
[leetcode 141. 环形链表](https://leetcode-cn.com/problems/linked-list-cycle/description/)




