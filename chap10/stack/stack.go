package stack

import (
	"errors"
)

type Stack[T any] struct {
	top   int
	size  int
	slice []T
}

func NewStack[T any](n int) *Stack[T] {
	return &Stack[T]{
		size:  n,
		slice: make([]T, n),
	}
}

func (s *Stack[T]) StackEmpty() bool {
	return s.top == 0
}

func (s *Stack[T]) Push(x T) error {
	if s.top == s.size {
		return errors.New("overflow")
	}
	s.top = s.top + 1
	s.slice[s.top] = x
	return nil
}

func (s *Stack[T]) Pop() (T, error) {
	var t T
	if s.StackEmpty() {
		return t, errors.New("underflow")
	}
	s.top = s.top - 1
	return s.slice[s.top+1], nil
}
