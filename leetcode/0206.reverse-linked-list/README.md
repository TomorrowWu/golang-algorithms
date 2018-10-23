
### 题目描述
反转一个单链表。

**示例:**
```
输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL
```

**进阶:**
```
你可以迭代或递归地反转链表。你能否用两种方法解决这道题？
```

### 解题思路
- 详见代码

### 代码实现

```Golang
// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var prev *ListNode
	cur := head
	for cur != nil {
		cur.Next, prev, cur = prev, cur, cur.Next
	}
	return prev
}
```

### GitHub
- [源码传送门](https://github.com/TomorrowWu/golang-algorithms/blob/master/leetcode/0206.reverse-linked-list/src/reverse-linked-list.go)
- 项目中会提供各种数据结构及算法的Golang实现, LeetCode解题思路及答案

###### 题目来源
[leetcode.206 反转链表](https://leetcode-cn.com/problems/reverse-linked-list/description/)
