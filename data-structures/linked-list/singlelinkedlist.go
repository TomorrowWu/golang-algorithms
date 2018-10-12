package main

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
