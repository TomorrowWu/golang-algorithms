package src

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
