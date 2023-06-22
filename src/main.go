package main

import (
	"fmt"

	datastructures "github.com/deankinane/algo-course/src/data_structures/ring_buffer"
)

func main() {
	rb := datastructures.NewRingBuffer(3)

	rb.Enqueue(1)
	rb.Enqueue(2)
	rb.Dequeue()
	rb.Enqueue(3)
	rb.Enqueue(4)
	rb.Enqueue(5)
	rb.Enqueue(6)

	fmt.Println(rb.ToString())
	fmt.Println(rb.ToArray())
}
