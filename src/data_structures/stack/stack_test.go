package datastructures

import "testing"

func TestStack(t *testing.T) {
	stack := new(Stack)

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)

	if v, _ := stack.Pop(); v != 4 {
		t.Errorf("Expected 4 but got %v", v)
	}

	if v, _ := stack.Pop(); v != 3 {
		t.Errorf("Expected 3 but got %v", v)
	}

	stack.Pop()
	stack.Pop()

	if _, err := stack.Pop(); err == nil {
		t.Error("Expected error but got none")
	}
}
