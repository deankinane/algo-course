package binarytree

import "github.com/deankinane/algo-course/src/data_structures/queue"

type BinaryNode[T comparable] struct {
	Val   T
	Left  *BinaryNode[T]
	Right *BinaryNode[T]
}

func (n *BinaryNode[T]) PreOrderTraverse() *[]T {
	data := []T{n.Val}

	data = append(data, walkPreOrder[T](n.Left)...)
	data = append(data, walkPreOrder[T](n.Right)...)

	return &data
}

func walkPreOrder[T comparable](node *BinaryNode[T]) []T {
	data := []T{}
	if node == nil {
		return data
	}

	data = append(data, node.Val)

	data = append(data, walkPreOrder[T](node.Left)...)
	data = append(data, walkPreOrder[T](node.Right)...)

	return data
}

func (node *BinaryNode[T]) PostOrderTraverse() []T {
	data := []T{}
	if node == nil {
		return data
	}

	data = append(data, walkPostOrder[T](node.Left)...)
	data = append(data, walkPostOrder[T](node.Right)...)

	return data
}

func walkPostOrder[T comparable](node *BinaryNode[T]) []T {
	data := []T{}
	if node == nil {
		return data
	}

	data = append(data, walkPreOrder[T](node.Left)...)
	data = append(data, walkPreOrder[T](node.Right)...)

	data = append(data, node.Val)

	return data
}

func (node *BinaryNode[T]) BreadthFirst() []T {
	q := queue.Queue[*BinaryNode[T]]{}
	data := []T{}

	q.Enqueue(node)
	for q.Length > 0 {
		v, _ := q.Dequeue()

		if v.Left != nil {
			q.Enqueue(v.Left)
		}
		if v.Right != nil {
			q.Enqueue(v.Right)
		}

		data = append(data, v.Val)
	}

	return data
}
