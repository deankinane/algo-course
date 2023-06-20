package datastructures

type LinkedListNode struct {
	Val  int
	Next *LinkedListNode
	Prev *LinkedListNode
}

func (n *LinkedListNode) Insert(newNode *LinkedListNode) {
	newNode.Next = n.Next

	if n.Next != nil {
		n.Next.Prev = newNode
	}

	n.Next = newNode
	newNode.Prev = n
}
