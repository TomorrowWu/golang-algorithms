package main

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	n := len(lists)
	switch n {
	case 0:
		return nil
	case 1:
		return lists[0]
	case 2:
		//针对两个链表进行归并排序
		return merge(lists[0], lists[1])
	default:
		key := n / 2
		//数组拆分,使下一次递归的lists的长度=2

		//优化思路: mergeKLists(lists[:key]),使用Goroutine+channel进行并发合并(归并排序的特点)
		return mergeKLists([]*ListNode{mergeKLists(lists[:key]), mergeKLists(lists[key:])})
	}

}

//merge 对两个有序链表进行归并排序
func merge(left *ListNode, right *ListNode) *ListNode {
	//head: 新的链表的head指针,保持不变
	//tail: 新链表的尾指针
	var head, tail *ListNode

	if left == nil {
		return right
	}

	if right == nil {
		return left
	}

	if left.Val < right.Val {
		head, tail, left = left, left, left.Next
	} else {
		head, tail, right = right, right, right.Next
	}

	//循环,直到某一个链表已遍历完
	for left != nil && right != nil {
		//找到下一个节点,添加到新链表的尾
		if left.Val < right.Val {
			tail.Next, left = left, left.Next
		} else {
			tail.Next, right = right, right.Next
		}
		//更新tail
		tail = tail.Next
	}

	//剩下的节点字节拼接到新链表尾部
	if left != nil {
		tail.Next = left
	}
	if right != nil {
		tail.Next = right
	}

	return head
}

func main() {

}
