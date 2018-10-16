package queue

import "testing"

func TestListQueue_EnQueue(t *testing.T) {
	q := NewLinkedListQueue()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)
	q.Enqueue(6)
	t.Log(q)
}

func TestListQueue_DeQueue(t *testing.T) {
	q := NewLinkedListQueue()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)
	q.Enqueue(6)
	t.Log(q)
	q.Dequeue()
	t.Log(q)
	q.Dequeue()
	t.Log(q)
	q.Dequeue()
	t.Log(q)
	q.Dequeue()
	t.Log(q)
	q.Dequeue()
	t.Log(q)
}
