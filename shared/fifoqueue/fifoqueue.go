package fifoqueue

type FifoQueue[T any] struct {
	items []T
	head  int
}

func New[T any](items ...T) FifoQueue[T] {
	return FifoQueue[T]{items: items}
}

func (q *FifoQueue[T]) IsEmpty() bool {
	return q.head == len(q.items)
}

func (q *FifoQueue[T]) Push(elem T) {
	q.items = append(q.items, elem)
}

func (q *FifoQueue[T]) Pull() T {
	if q.IsEmpty() {
		return *new(T)
	}
	elem := q.items[q.head]
	q.head++
	if q.head > 1_000_000 {
		q.items = q.items[q.head:]
		q.head = 0
	}
	return elem
}
