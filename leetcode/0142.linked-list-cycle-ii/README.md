### 题目描述
给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。

**说明**：不允许修改给定的链表。
**进阶:**
```
你能否不使用额外空间解决此题？
```

### 解题思路
- 无环链表，最后一个节点为nil，有环链表可以无限循环next下去
- 不用额外空间：快慢节点，慢节点一次走一步，快节点一次走两步，当进入环中，每次循环，快节点会离慢节点近一步，快节点最终会追上慢节点
- 用额外空间： 用map存走过的节点，第一个走过的节点就是环的入口
- 不用额外空间
![环形链表](https://github.com/TomorrowWu/pictures/blob/master/algorithms/%E7%8E%AF%E5%BD%A2%E9%93%BE%E8%A1%A8.jpg?raw=true)
```
设：链表头是X，环的第一个节点是Y，slow和fast第一次的交点是Z。各段的长度分别是a,b,c，如图所示
第一次相遇时slow走过的距离：a+b，fast走过的距离：a+b+c+b
因为fast的速度是slow的两倍，所以fast走的距离是slow的两倍，有 2(a+b) = a+b+c+b，可以得到a=c（这个结论很重要！）
这时候，slow从X出发，fast从Z出发，以相同速度走，同时到达Y，Y就是环的入口，即第一个节点
```

### 代码实现
```Golang
// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head != nil {
		slow := head
		fast := head
		for fast != nil && fast.Next != nil {
			slow = slow.Next
			fast = fast.Next.Next
			if slow == fast {
				// X---------Y---------Z
				//x起点，y环的起点，z是相遇点
				//链表头是X，环的第一个节点是Y，slow和fast第一次的交点是Z。各段的长度分别是a,b,c
				//相遇时，fast走过的距离为 2(a+b)=a+b+c+b，得到距离a=c
				//则 slow从head开始走，fast从Z点开始走，速度都是一次走一步，slow和fast同时到达Y点，即环的入口
				slow = head
				for slow != fast {
					slow = slow.Next
					fast = fast.Next
				}
				return slow
			}
		}
	}
	return nil
}
```

### GitHub
- [源码传送门](https://github.com/TomorrowWu/golang-algorithms/blob/master/leetcode/0142.linked-list-cycle-ii/src/linked-list-cycle-ii.go)
- 项目中会提供各种数据结构及算法的Golang实现, LeetCode解题思路及答案

###### 题目来源
[leetcode 142. 环形链表 II](https://leetcode-cn.com/problems/linked-list-cycle-ii/description/)




