package queue

type Queue[T any] struct {
	head  int
	tail  int
	size  int
	slice []T
}

func NewQueue[T any](size int) *Queue[T] {
	return &Queue[T]{
		size:  size,
		slice: make([]T, size),
	}
}

func (q *Queue[T]) EnQueue(x T) error {
	q.slice[q.tail] = x
	q.tail = q.tail + 1
	if q.tail == q.size {
		q.tail = 0
		return nil
	}
	return nil
}

func (q *Queue[T]) DeQueue() (T, error) {
	x := q.slice[q.head]
	q.head = q.head + 1
	if q.head == q.size {
		q.head = 0
		return x, nil
	}
	return x, nil
}
