package linkedlist

import "fmt"

// ListNode defines a node of single-linked-list
type ListNode struct {
	next  *ListNode
	value interface{}
}

// LinkedList defines a linked-list
type LinkedList struct {
	head   *ListNode
	length int
}

// NewListNode created a node
func NewListNode(v interface{}) *ListNode {
	return &ListNode{nil, v}
}

// GetNext returns next of node
func (n *ListNode) GetNext() *ListNode {
	return n.next
}

// GetValue returns value of node
func (n *ListNode) GetValue() interface{} {
	return n.value
}

// NewLinkedList creates a empty linked-list
func NewLinkedList() *LinkedList {
	//head not store value
	return &LinkedList{
		head:   NewListNode(0),
		length: 0,
	}
}

// InsertAfter inserts new node after node p
func (l *LinkedList) InsertAfter(p *ListNode, v interface{}) bool {
	if p == nil {
		return false
	}

	newNode := NewListNode(v)
	oldNext := p.next
	p.next = newNode
	newNode.next = oldNext
	l.length++
	return true
}

// InsertBefore inserts new node before node p
func (l *LinkedList) InsertBefore(p *ListNode, v interface{}) bool {
	if p == nil || p == l.head {
		return false
	}

	cur := l.head.next
	pre := l.head
	for cur != nil {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.next
	}

	if cur == nil {
		return false
	}

	newNode := NewListNode(v)
	pre.next = newNode
	newNode.next = cur
	l.length++
	return true
}

// InsertToHead inserts new node after head
func (l *LinkedList) InsertToHead(v interface{}) bool {
	return l.InsertAfter(l.head, v)
}

// InsertToTail inserts new node to tail
func (l *LinkedList) InsertToTail(v interface{}) bool {
	cur := l.head
	for cur.next != nil {
		cur = cur.next
	}

	return l.InsertAfter(cur, v)
}

// FindByIndex returns node by index
func (l *LinkedList) FindByIndex(index int) *ListNode {
	if index >= l.length {
		return nil
	}

	//head not store value
	cur := l.head.next
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	return cur
}

// DeleteNode deletes node
func (l *LinkedList) DeleteNode(p *ListNode) bool {
	if nil == p {
		return false
	}
	cur := l.head.next
	pre := l.head
	for nil != cur {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.next
	}
	if nil == cur {
		return false
	}
	pre.next = p.next
	p = nil
	l.length--
	return true
}

// Print prints linked-list
func (l *LinkedList) Print() {
	cur := l.head.next
	format := ""
	for nil != cur {
		format += fmt.Sprintf("%+v", cur.GetValue())
		cur = cur.next
		if nil != cur {
			format += "->"
		}
	}
	fmt.Println(format)
}

// Reverse reverses the linked-list.
//
// Time complexity is O(n)
func (l *LinkedList) Reverse() {
	if l.head == nil || l.head.next == nil || l.head.next.next == nil {
		return
	}

	var pre *ListNode
	cur := l.head.next
	for cur != nil {
		tmp := cur.next
		cur.next = pre
		pre = cur
		cur = tmp
	}
	l.head.next = pre
}

// HasCycle judges linked-list if it is cycled-list
func (l *LinkedList) HasCycle() bool {
	if l.head != nil {
		slow := l.head
		fast := l.head
		for fast != nil && fast.next != nil {
			slow = slow.next
			fast = fast.next.next
			if slow == fast {
				return true
			}
		}
	}
	return false
}

// MergeSortedList merges two sorted list
func MergeSortedList(l1, l2 *LinkedList) *LinkedList {
	if l1 == nil || l1.head == nil || l1.head.next == nil {
		return l2
	}

	if l2 == nil || l2.head == nil || l2.head.next == nil {
		return l1
	}

	newList := &LinkedList{
		head: &ListNode{
			next:  nil,
			value: nil,
		},
		length: 0,
	}

	cur := newList.head
	cur1 := l1.head.next
	cur2 := l2.head.next
	for cur1 != nil && cur2 != nil {
		if cur1.value.(int) <= cur2.value.(int) {
			cur.next = cur1
			cur1 = cur1.next
		} else {
			cur.next = cur2
			cur2 = cur2.next
		}
		cur = cur.next
	}

	if cur1 != nil {
		cur.next = cur1
	} else if cur2 != nil {
		cur.next = cur2
	}

	return newList

}

// DeleteBottomN deletes The bottom NTH node
func (l *LinkedList) DeleteBottomN(n int) {
	if n <= 0 || l.head == nil || l.head.next == nil {
		return
	}

	fast := l.head
	for i := 1; i <= n && fast != nil; i++ {
		fast = fast.next
	}

	if fast == nil {
		return
	}

	// \-----n(fast)--------n-------\
	// \-----n--------n(slow)-------\(fast)
	slow := l.head
	for fast.next != nil {
		slow = slow.next
		fast = fast.next
	}
	slow.next = slow.next.next
}

// FindMiddleNode returns middle node
func (l *LinkedList) FindMiddleNode() *ListNode {
	if l.head == nil || l.head.next == nil {
		return nil
	}
	if l.head.next.next == nil {
		return l.head.next
	}

	slow, fast := l.head, l.head
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	return slow
}
