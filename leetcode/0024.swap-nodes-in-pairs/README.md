### 题目描述
给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。

**示例:**
```
给定 1->2->3->4, 你应该返回 2->1->4->3.
```
**说明:**
- 你的算法只能使用常数的额外空间。
- **你不能只是单纯的改变节点内部的值**，而是需要实际的进行节点交换。

### 代码实现
```Golang
// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var prev *ListNode
	cur := head
	head = cur.Next
	for ; cur != nil && cur.Next != nil; cur = cur.Next {
		next := cur.Next
		//注意：第一次循环时，prev为nil
		if prev != nil {
			prev.Next = next
		}
		//交换两个节点
		cur.Next, next.Next, prev = next.Next, cur, cur
	}

	return head
}
```

### GitHub
- [源码传送门](https://github.com/TomorrowWu/golang-algorithms/blob/master/leetcode/0024.swap-nodes-in-pairs/src/swap-nodes-in-pairs.go)
- 项目中会提供各种数据结构及算法的Golang实现, LeetCode解题思路及答案

###### 题目来源
[leetcode 24. 两两交换链表中的节点](https://leetcode-cn.com/problems/swap-nodes-in-pairs/description/)




