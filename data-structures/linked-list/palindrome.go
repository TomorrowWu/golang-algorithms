package linkedlist

/**
思路1:用一个切片存前半部分
时间复杂度：O(N)
空间复杂度：O(N)
*/
func isPalindrome1(l *LinkedList) bool {
	if l.length <= 1 {
		return true
	}

	prevHalf := make([]interface{}, 0, l.length/2)
	cur := l.head
	for i := 1; i <= l.length; i++ {
		cur = cur.next
		if l.length%2 != 0 && i == l.length/2+1 {
			//如果链表有奇数个节点，中间的直接忽略
			continue
		}
		if i <= l.length/2 {
			//前一半存储
			prevHalf = append(prevHalf, cur.GetValue())
		} else {
			//后一半与前一半进行对比
			if prevHalf[l.length-i] != cur.GetValue() {
				return false
			}
		}
	}

	return true
}

/*
思路2
找到链表中间节点，将前半部分转置，再从中间向左右遍历对比
使用的常数个指针变量,比对
时间复杂度：O(N)
空间复杂度：O(1)
*/
func isPalindrome2(l *LinkedList) bool {
	if l.length <= 1 {
		return true
	}

	isPalindrome := true
	step := l.length / 2
	var pre *ListNode
	cur := l.head.next
	next := l.head.next.next
	for i := 1; i <= step; i++ {
		tmp := cur.GetNext()
		cur.next = pre
		pre = cur
		cur = tmp
		next = cur.GetNext()
	}
	mid := cur

	var left, right *ListNode = pre, nil
	if l.length%2 != 0 {
		right = mid.GetNext()
	} else {
		right = mid
	}

	for left != nil && right != nil {
		if left.GetValue() != right.GetValue() {
			isPalindrome = false
			break
		}
		left = left.GetNext()
		right = right.GetNext()
	}

	//复原链表
	cur = pre
	pre = mid
	for cur != nil {
		next = cur.GetNext()
		cur.next = pre
		pre = cur
		cur = next
	}

	return isPalindrome
}
