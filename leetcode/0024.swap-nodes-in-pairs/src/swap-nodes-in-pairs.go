package src

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
