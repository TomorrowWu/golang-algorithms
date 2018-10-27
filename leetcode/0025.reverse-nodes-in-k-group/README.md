### 题目描述
给出一个链表，每 k 个节点一组进行翻转，并返回翻转后的链表。

k 是一个正整数，它的值小于或等于链表的长度。如果节点总数不是 k 的整数倍，那么将最后剩余节点保持原有顺序。

**示例:**
```
给定这个链表：1->2->3->4->5

当 k = 2 时，应当返回: 2->1->4->3->5

当 k = 3 时，应当返回: 3->2->1->4->5
```
**说明:**
- 你的算法只能使用常数的额外空间。
- **你不能只是单纯的改变节点内部的值**，而是需要实际的进行节点交换。

### 解题思路
```
1. 取链表的前K个节点，如果够K个节点，就截断后进行反转，不够K个节点，说明处理完了，return
2. 反转完前K个节点后，使用递归，处理后面的链表
```

### 代码实现
```Golang
// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k < 2 || head == nil || head.Next == nil {
		return head
	}

	tail, needReverse := getTail(head, k)

	if needReverse {
		tailNext := tail.Next
		//斩断 tail 后的链接
		tail.Next = nil
		head, tail = reverse(head, tail)

		//tail 后面接上尾部的递归处理
		tail.Next = reverseKGroup(tailNext, k)
	}

	return head
}

func getTail(head *ListNode, k int) (*ListNode, bool) {
	for k > 1 && head != nil {
		head = head.Next
		k--
	}
	return head, k == 1 && head != nil
}

func reverse(head, tail *ListNode) (*ListNode, *ListNode) {
	//反转链表
	prev, cur := head, head.Next
	for cur != nil {
		prev, cur, cur.Next = cur, cur.Next, prev
	}
	return tail, head
}
```

### GitHub
- [源码传送门](https://github.com/TomorrowWu/golang-algorithms/blob/master/leetcode/0025.reverse-nodes-in-k-group/src/reverse-nodes-in-k-group.go)
- 项目中会提供各种数据结构及算法的Golang实现, LeetCode解题思路及答案

###### 题目来源
[leetcode 25. k个一组翻转链表](https://leetcode-cn.com/problems/reverse-nodes-in-k-group/description/)




