[0023_合并K个排序链表](https://leetcode-cn.com/problems/merge-k-sorted-lists/description/)

## 题目描述
合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度。

示例:
```
输入:
[
  1->4->5,
  1->3->4,
  2->6
]
输出: 1->1->2->3->4->4->5->6
```

## 算法

```Golang
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
```

## 个人思路

```
1. 对已经有序的多个链表进行合并,可以借鉴归并排序,分治法的思想,层层递归
2. 两个链表可以进行遍历比较节点大小,合并为一个新的链表
```

## 总结
- 分治法的特点:各层分治递归可以同时进行,优化思路可以采用Goroutine+channel,在此笔者就不进行优化了,不是本文重点

#### 递归优缺点
优点:
```
1. 代码简洁
```

缺点:
```
1. 空间消耗大,每一次函数调用都需要在内存栈中分配空间保存参数,返回地址以及临时变量
2. 栈里面压入和弹出都需要时间
3. 递归有栈溢出的问题
```

#### 循环
优点
```
1. 和递归相比,空间消耗小
```
缺点
```
1. 代码可读性不如递归
```

## GitHub
- [项目源码在这里](https://github.com/TomorrowWu/dataStructures-algorithm/tree/master/leetcode)
- 笔者会一直维护该项目,对leetcode中的算法题进行解决,并写下自己的思路和见解,致力于人人都能看懂的算法

## 个人公众号
- 喜欢的朋友可以关注,谢谢支持
- 记录90后码农的学习与生活

![image](https://upload-images.jianshu.io/upload_images/5815624-4a8b49cfbaf037dd.jpg?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)
