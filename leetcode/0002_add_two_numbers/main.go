package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

//++++++++++++++++++++++++++++++++++++易理解的方式1(转化为数据存,直观,效率低)+++++++++++++++++++++++++++++++++++++++++++

func addNode(head *ListNode, node *ListNode) *ListNode {
	temp := head
	for {
		if head == nil {
			head = node
			return head
		}
		//找到最后一个节点
		if temp.Next == nil {
			break
		}
		temp = temp.Next
	}
	temp.Next = node
	return head
}

func show(head *ListNode) {
	temp := head
	for {
		if temp == nil {
			break
		}
		fmt.Printf("%d", temp.Val)
		temp = temp.Next
		if temp != nil {
			fmt.Printf("-->")
		}
	}
}

//addTwoNumbers 将每一位上的和存储
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	arr1 := listNodeToArr(l1)
	arr2 := listNodeToArr(l2)

	var maxArr []int
	var minArr []int
	if len(arr1) <= len(arr2) {
		minArr = arr1
		maxArr = arr2
	} else {
		minArr = arr2
		maxArr = arr1
	}

	maxLen := len(maxArr)
	minLen := len(minArr)
	resultArr := make([]int, maxLen+1, maxLen+1) //可能最后一位的和大于10,要进1,这里预留一位

	var head *ListNode
	for i, _ := range maxArr {
		sum := 0
		if i <= minLen-1 {
			sum = minArr[i] + maxArr[i]
		} else {
			sum = maxArr[i]
		}

		resultArr[i] += sum
		if resultArr[i] >= 10 {
			//10进制加1
			resultArr[i+1] += 1
			resultArr[i] -= 10
		}

		head = addNode(head, &ListNode{
			Val:  resultArr[i],
			Next: nil,
		})
	}
	//最高位node
	num := resultArr[len(resultArr)-1]
	if num != 0 {
		//最高位是冗余存的,为0时,只是占位符
		head = addNode(head, &ListNode{
			Val:  num,
			Next: nil,
		})
	}
	return head
}

func listNodeToArr(l *ListNode) []int {
	var arr []int
	temp := l
	for {
		if temp == nil {
			break
		}
		arr = append(arr, temp.Val)
		temp = temp.Next
	}
	return arr
}

//++++++++++++++++++++++++++++++++++++易理解的方式1+++++++++++++++++++++++++++++++++++++++++++

//不转化为数组,效率优化版本(参考官方答案)
//使用数组,循环赋值耗时长
func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	//只是链表头,不存值
	head := &ListNode{
		Val:  0,
		Next: nil,
	}

	curNode := head
	carry := 0 //存上一位的和后,满10进1的值
	p, q := l1, l2
	for p != nil || q != nil {
		x, y := 0, 0
		if p != nil {
			x = p.Val
		}
		if q != nil {
			y = q.Val
		}

		sum := carry + x + y
		carry = sum / 10 //即进位
		curNode.Next = &ListNode{
			Val:  sum % 10, //该位的值
			Next: nil,
		}
		curNode = curNode.Next

		if p != nil {
			p = p.Next
		}
		if q != nil {
			q = q.Next
		}
	}

	//最后一位
	if carry > 0 {
		curNode.Next = &ListNode{
			Val:  carry,
			Next: nil,
		}
	}

	return head.Next
}

func main() {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	sumNode := addTwoNumbers(l1, l2)
	show(sumNode)
}
