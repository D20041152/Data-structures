package main

import "fmt"

type Stack struct {
	stack []int
	Top   int
}

func NewStack() *Stack {
	return &Stack{stack: []int{}, Top: 0}
}

func (s *Stack) Push(a int) {
	s.stack = append(s.stack, a)
	s.Top = a
}

func (s *Stack) Pop() (int, error) {
	N := len(s.stack)
	if N < 1 {
		return -1000, fmt.Errorf("Stack in empty")
	}
	elem := s.stack[N-1]
	s.stack = s.stack[:N-1]
	if N != 1 {
		s.Top = s.stack[N-2]
	} else {
		s.Top = 0
	}
	return elem, nil
}

func (s *Stack) IsEmpty() bool {
	if len(s.stack) == 0 {
		return true
	} else {
		return false
	}
}
