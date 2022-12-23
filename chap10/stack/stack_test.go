package stack_test

import (
	"testing"

	"github.com/GRTheory/algorithms/chap10/stack"
)

func TestStack(t *testing.T) {
	myStack := stack.NewStack[string](10)
	myStack.Push("a")
	myStack.Push("b")
	myStack.Push("c")

	result, _ := myStack.Pop()
	t.Log(result)
	result, _ = myStack.Pop()
	t.Log(result)
	result, _ = myStack.Pop()
	t.Log(result)
}
