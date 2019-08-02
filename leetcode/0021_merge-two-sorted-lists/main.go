package main

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	n := &ListNode{
		Val:  0,
		Next: nil,
	}
	cur, c1, c2 := n, l1, l2
	// 直到某一链结束
	for c1 != nil && c2 != nil {
		if c1.Val < c2.Val {
			cur.Next = &ListNode{
				Val:  c1.Val,
				Next: nil,
			}
			c1 = c1.Next

		} else {
			cur.Next = &ListNode{
				Val:  c2.Val,
				Next: nil,
			}
			c2 = c2.Next
		}
		cur = cur.Next
	}
	if c1 == nil {
		cur.Next = c2
	}
	if c2 == nil {
		cur.Next = c1
	}
	return n.Next
}
