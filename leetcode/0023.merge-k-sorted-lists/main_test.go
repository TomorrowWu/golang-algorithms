package main

import (
	"fmt"
	"testing"
)

func Test_mergeKLists(t *testing.T) {
	list1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  5,
				Next: nil,
			},
		},
	}
	list2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	list3 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val:  6,
			Next: nil,
		},
	}
	kLists := mergeKLists([]*ListNode{list1, list2, list3})
	showListNode(kLists)
}

func showListNode(listNode *ListNode) {
	tmp := listNode
	for tmp != nil {
		fmt.Printf("%d->", tmp.Val)
		tmp = tmp.Next
	}
}
