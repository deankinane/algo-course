package datastructures

import (
	"fmt"
	"testing"
)

func TestRingBuffer(t *testing.T) {
	rb := NewRingBuffer(3)

	rb.Enqueue(1)
	rb.Enqueue(2)

	if val, _ := rb.Dequeue(); val != 1 {
		t.Error("Expected dequeue value of 1, got ", val)
	}

	rb.Enqueue(3)
	rb.Enqueue(4)

	if test := rb.ToArray(); len(test) != 3 {
		t.Error("Expteced array to have 3 elements, got ", len(test))
	} else if fmt.Sprint(test) != "[2 3 4]" {

		t.Error("Exptected array to be [2 3 4], got ", test)
	}

	rb.Enqueue(5)
	rb.Enqueue(6)

	if test := rb.ToArray(); len(test) != 5 {
		t.Error("Expteced array to have 5 elements, got ", len(test))
	} else if fmt.Sprint(test) != "[2 3 4 5 6]" {

		t.Error("Exptected array to be [2 3 4 5 6], got ", test)
	}

	rb.Enqueue(7)
	rb.Enqueue(8)

	if val, _ := rb.Dequeue(); val != 2 {
		t.Error("Expected dequeue value of 2, got ", val)
	}

	rb.Dequeue()
	rb.Enqueue(9)
	rb.Enqueue(10)

	if test := rb.ToArray(); len(test) != 7 {
		t.Error("Expteced array to have 7 elements, got ", len(test))
	} else if fmt.Sprint(test) != "[4 5 6 7 8 9 10]" {

		t.Error("Exptected array to be [4 5 6 7 8 9 10], got ", test)
	}
}
