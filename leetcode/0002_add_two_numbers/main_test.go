package main

import "testing"

func Test_addTwoNumbers(t *testing.T) {
	l1 := &ListNode{
		Val:  1,
		Next: nil,
	}
	l2 := &ListNode{
		Val: 9,
		Next: &ListNode{
			Val:  9,
			Next: nil,
		},
	}
	sumNode := addTwoNumbers(l1, l2)
	show(sumNode)
}

func Test_addTwoNumbers2(t *testing.T) {
	l1 := &ListNode{
		Val:  1,
		Next: nil,
	}
	l2 := &ListNode{
		Val: 9,
		Next: &ListNode{
			Val:  9,
			Next: nil,
		},
	}
	sumNode := addTwoNumbers2(l1, l2)
	show(sumNode)
}
