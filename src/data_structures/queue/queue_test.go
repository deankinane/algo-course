package datastructures

import (
	"testing"
)

func TestQueue(t *testing.T) {
	q := new(Queue)

	if _, err := q.Peek(); err == nil {
		t.Error("Expected empty queue error")
	}

	for i := 1; i <= 5; i++ {
		q.Enqueue(i)
	}

	if q.Length != 5 {
		t.Errorf("Expected length to be 5 but got %v", q.Length)
	}

	if v, err := q.Dequeue(); err != nil {
		t.Error("Expected 1 but got error")
	} else if v != 1 {
		t.Errorf("Expected 1 but got %v", v)
	}

	if v, err := q.Dequeue(); err != nil {
		t.Error("Expected 2 but got error")
	} else if v != 2 {
		t.Errorf("Expected 2 but got %v", v)
	}

	if q.Tail.Val != 5 {
		t.Errorf("Expteded tail to be 5 but got %v", q.Tail.Val)
	}

	if v, err := q.Peek(); err != nil {
		t.Error("Expected 3 but got error")
	} else if v != 3 {
		t.Errorf("Expected 3 but got %v", v)
	}

	q.Dequeue()
	q.Dequeue()
	q.Dequeue()

	if _, err := q.Dequeue(); err == nil {
		t.Error("Expected empty queue error")
	}
}
