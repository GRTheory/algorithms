package queue_test

import (
	"testing"

	"github.com/GRTheory/algorithms/chap10/queue"
)

func TestQueue(t *testing.T) {
	myQueue := queue.NewQueue[string](3)
	myQueue.EnQueue("1")
	myQueue.EnQueue("2")
	myQueue.EnQueue("3")
	result, _ := myQueue.DeQueue()
	t.Log(result)
	myQueue.EnQueue("4")
	result, _ = myQueue.DeQueue()
	t.Log(result)
}