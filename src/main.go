package main

import (
	"fmt"

	datastructures "github.com/deankinane/algo-course/src/data_structures/linked_list"
)

func main() {
	head := new(datastructures.LinkedListNode)
	head.Val = 1

	current := head
	for i := 2; i <= 10; i++ {
		newNode := new(datastructures.LinkedListNode)
		newNode.Val = i
		current.Insert(newNode)
		current = newNode
	}

	current = head
	last := head
	for current != nil {
		fmt.Println(current.Val)
		last = current
		current = current.Next
	}

	for last != nil {
		fmt.Println(last.Val)
		last = last.Prev
	}
}
