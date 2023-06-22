package datastructures

import (
	"errors"
	"fmt"
)

type RingBuffer struct {
	array  []int
	tail   int
	head   int
	length int
}

func NewRingBuffer(length int) RingBuffer {
	return RingBuffer{
		array:  make([]int, length),
		tail:   -1,
		head:   -1,
		length: length,
	}
}

func (rb *RingBuffer) Enqueue(val int) {
	if rb.head == -1 {
		rb.head = 0
		rb.tail = 0
		rb.array[0] = val
		return
	}

	if (rb.tail+1)%rb.length == rb.head%rb.length {
		copy := rb.ToArray()
		rb.tail = len(copy) - 1
		rb.head = 0

		copy = append(copy, make([]int, rb.length)...)
		rb.length *= 2 // exponential growth is a dumb idea but fine for this simple implementation
		rb.array = copy
	}

	rb.tail++
	rb.array[rb.tail%rb.length] = val
}

func (rb *RingBuffer) Dequeue() (int, error) {
	if rb.head == -1 {
		return 0, errors.New("empty list")
	}

	val := rb.array[rb.head%rb.length]
	rb.head++

	return val, nil
}

func (rb *RingBuffer) ToString() string {
	return fmt.Sprintf("Head: %v\nTail: %v\nArray: %v", rb.head, rb.tail, rb.array)
}

func (rb *RingBuffer) ToArray() []int {
	order := []int{}

	for i := 0; i < len(rb.array); i++ {
		cur := (rb.head + i) % rb.length
		if cur > rb.tail {
			break
		}
		order = append(order, rb.array[cur])
	}

	return order
}
