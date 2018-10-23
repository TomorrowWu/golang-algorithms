package src

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 解法1
// 用数组存前面的一半节点的值
// 时间复杂度：O(N)
// 空间复杂度：O(N)
func isPalindrome(head *ListNode) bool {
	// 空链表,算回文
	if head == nil {
		return true
	}

	var data []int
	for cur := head; cur != nil; cur = cur.Next {
		data = append(data, cur.Val)
	}

	for i, j := 0, len(data)-1; i <= j; {
		if data[i] != data[j] {
			return false
		}
		i++
		j--
	}
	return true
}

// 解法2
// 找到链表中间节点，将前半部分转置，再从中间向左右遍历对比
// 时间复杂度：O(N)
// 空间复杂度：O(1)
func isPalindrome2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	isPalindrome := true

	//链表长度
	length := 0
	for cur := head; cur != nil; cur = cur.Next {
		length++
	}

	//将前半部分反转
	step := length / 2
	var prev *ListNode
	cur := head
	for i := 1; i <= step; i++ {
		cur.Next, prev, cur = prev, cur, cur.Next
	}
	mid := cur

	var left, right *ListNode = prev, nil
	if length%2 == 0 {
		//长度为偶数
		right = mid
	} else {
		right = mid.Next
	}

	//从中间向左右两边遍历对比
	for left != nil && right != nil {
		if left.Val != right.Val {
			//值不相等,不是回文链表
			isPalindrome = false
			break
		}
		left = left.Next
		right = right.Next
	}

	//将前半部分反转的链表进行复原
	cur = prev
	prev = mid
	for cur != nil {
		cur.Next, prev, cur = prev, cur, cur.Next
	}

	return isPalindrome
}
