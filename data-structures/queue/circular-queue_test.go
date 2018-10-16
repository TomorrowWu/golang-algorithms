package queue

import "testing"

func TestCircularQueue_EnQueue(t *testing.T) {
	q := NewCircularQueue(5)
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)
	q.Enqueue(6)
	t.Log(q)
}

func TestCircularQueue_DeQueue(t *testing.T) {
	q := NewCircularQueue(5)
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)
	q.Enqueue(6)
	t.Log(q)
	t.Log(q.Dequeue())
	t.Log(q)
	q.Enqueue(5)
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
