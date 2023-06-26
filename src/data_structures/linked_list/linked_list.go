package datastructures

import "errors"

type LinkedListNode[T comparable] struct {
	Item T
	Next *LinkedListNode[T]
	Prev *LinkedListNode[T]
}

type LinkedList[T comparable] struct {
	Head   *LinkedListNode[T]
	Tail   *LinkedListNode[T]
	Length int
}

func (l *LinkedList[T]) Prepend(item T) {
	node := LinkedListNode[T]{Item: item}
	l.Length++
	if l.Head == nil {
		l.Head = &node
		l.Tail = &node
		return
	}

	node.Next = l.Head
	l.Head.Prev = &node
	l.Head = &node
}

func (l *LinkedList[T]) Append(item T) {
	node := LinkedListNode[T]{Item: item}
	l.Length++
	if l.Tail == nil {
		l.Head = &node
		l.Tail = &node
		return
	}

	node.Prev = l.Tail
	l.Tail.Next = &node
	l.Tail = &node
}

func (l *LinkedList[T]) InsertAt(item T, index int) {
	node := LinkedListNode[T]{Item: item}

	cur := l.Head
	for i := 0; i < index; i++ {
		cur = cur.Next
	}

	node.Next = cur
	node.Prev = cur.Prev
	cur.Prev.Next = &node
}

func (l *LinkedList[T]) Remove(item T) (T, error) {
	cur := l.Head
	for cur != nil {
		if cur.Item == item {
			cur.Prev.Next = cur.Next
			cur.Next.Prev = cur.Prev
			return item, nil
		}

		cur = cur.Next
	}
	return *new(T), errors.New("not found in list")
}

func (l *LinkedList[T]) RemoveAt(index int) (T, error) {
	cur := l.Head
	for i := 0; i < index && cur != nil; i++ {
		cur = cur.Next
	}

	if cur != nil {
		cur.Prev.Next = cur.Next
		cur.Next.Prev = cur.Prev
		return cur.Item, nil
	}

	return *new(T), errors.New("not found in list")
}

func (l *LinkedList[T]) Get(index int) (T, error) {
	if index > l.Length {
		return *new(T), errors.New("index not in list")
	}

	cur := l.Head
	for i := 0; i < index && cur != nil; i++ {
		cur = cur.Next
	}

	if cur == nil {
		return *new(T), errors.New("index not in list")
	}

	return cur.Item, nil
}
