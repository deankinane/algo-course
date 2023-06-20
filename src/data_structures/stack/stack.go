package datastructures

import "errors"

type StackNode struct {
	Val  int
	Next *StackNode
}

type Stack struct {
	Lenght int
	Head   *StackNode
}

func (s *Stack) Push(val int) {
	node := new(StackNode)
	node.Val = val
	s.Lenght++

	if s.Head == nil {
		s.Head = node
		return
	}

	node.Next = s.Head
	s.Head = node
}

func (s *Stack) Pop() (int, error) {
	if s.Head == nil {
		return 0, errors.New("empty stack")
	}

	s.Lenght--
	val := s.Head.Val
	s.Head = s.Head.Next
	return val, nil
}

func (s *Stack) Peek() (int, error) {
	if s.Head == nil {
		return 0, errors.New("empty stack")
	}

	return s.Head.Val, nil
}
