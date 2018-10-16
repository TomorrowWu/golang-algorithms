package queue

import "fmt"

// ListNode defines a node of linked-list
type ListNode struct {
	val  interface{}
	next *ListNode
}

// LinkedListQueue defines a queue based on linked-list
type LinkedListQueue struct {
	head   *ListNode
	tail   *ListNode
	length int
}

// NewLinkedListQueue creates a queue
func NewLinkedListQueue() *LinkedListQueue {
	return &LinkedListQueue{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}

// Enqueue puts a element in the tail of queue
func (q *LinkedListQueue) Enqueue(v interface{}) {
	node := &ListNode{
		val:  v,
		next: nil,
	}
	if q.tail == nil {
		q.tail = node
		q.head = node
	} else {
		q.tail.next = node
		q.tail = node
	}
	q.length++
}

// Dequeue fetches a element from queue
func (q *LinkedListQueue) Dequeue() interface{} {
	if q.head == nil {
		return nil
	}
	v := q.head.val
	q.head = q.head.next
	q.length--
	return v
}

// String returns object string
func (q *LinkedListQueue) String() string {
	if q.head == nil {
		return "empty queue"
	}
	result := "head"
	for cur := q.head; cur != nil; cur = cur.next {
		result += fmt.Sprintf("<-%+v", cur.val)
	}
	result += "<-tail"
	return result
}
