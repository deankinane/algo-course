package datastructures

import "errors"

type QueueNode struct {
	Val  int
	next *QueueNode
}

type Queue struct {
	head *QueueNode
	tail *QueueNode
}

func (q *Queue) Enqueue(val int) {
	node := new(QueueNode)
	node.Val = val

	if q.head == nil {
		q.head = node
		q.tail = node
		return
	}

	q.tail.next = node
	q.tail = node
}

func (q *Queue) Dequeue() (int, error) {
	if q.head == nil {
		return 0, errors.New("empty queue")
	}

	val := q.head.Val
	q.head = q.head.next
	return val, nil
}

func (q *Queue) Peek() (int, error) {
	if q.head == nil {
		return 0, errors.New("empty queue")
	}
	return q.head.Val, nil
}
