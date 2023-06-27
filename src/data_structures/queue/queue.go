package queue

import "errors"

type QueueNode[T comparable] struct {
	Item T
	Next *QueueNode[T]
}

type Queue[T comparable] struct {
	Length int
	Head   *QueueNode[T]
	Tail   *QueueNode[T]
}

func (q *Queue[T]) Enqueue(val T) {
	node := new(QueueNode[T])
	node.Item = val
	q.Length++

	if q.Head == nil {
		q.Head = node
		q.Tail = node
		return
	}

	q.Tail.Next = node
	q.Tail = node
}

func (q *Queue[T]) Dequeue() (T, error) {
	if q.Head == nil {
		return *new(T), errors.New("empty queue")
	}

	val := q.Head.Item
	q.Head = q.Head.Next
	q.Length--

	return val, nil
}

func (q *Queue[T]) Peek() (T, error) {
	if q.Head == nil {
		return *new(T), errors.New("empty queue")
	}
	return q.Head.Item, nil
}
