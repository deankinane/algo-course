package datastructures

import "errors"

type QueueNode struct {
	Val  int
	Next *QueueNode
}

type Queue struct {
	Length int
	Head   *QueueNode
	Tail   *QueueNode
}

func (q *Queue) Enqueue(val int) {
	node := new(QueueNode)
	node.Val = val
	q.Length++

	if q.Head == nil {
		q.Head = node
		q.Tail = node
		return
	}

	q.Tail.Next = node
	q.Tail = node
}

func (q *Queue) Dequeue() (int, error) {
	if q.Head == nil {
		return 0, errors.New("empty queue")
	}

	val := q.Head.Val
	q.Head = q.Head.Next
	q.Length--

	return val, nil
}

func (q *Queue) Peek() (int, error) {
	if q.Head == nil {
		return 0, errors.New("empty queue")
	}
	return q.Head.Val, nil
}
