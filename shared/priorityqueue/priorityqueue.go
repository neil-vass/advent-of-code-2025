package priorityqueue

import "container/heap"

type PriorityQueue[T any] struct{ q *internalQueue[T] }

func New[T any]() PriorityQueue[T] {
	q := make(internalQueue[T], 0)
	return PriorityQueue[T]{&q}
}

func (pq PriorityQueue[T]) Push(value T, priority float64) {
	heap.Push(pq.q, &internalItem[T]{value, priority})
}

func (pq PriorityQueue[T]) Pull() T {
	if pq.IsEmpty() {
		return *new(T)
	}
	item := heap.Pop(pq.q).(*internalItem[T])
	return item.value
}

func (pq PriorityQueue[T]) IsEmpty() bool {
	return pq.q.Len() == 0
}

// internalItem and internalQueue use example code from the Go documentation:
// https://pkg.go.dev/container/heap@go1.25.4#example-package-PriorityQueue
// with some adaptations (no update(), so no need for index tracking, and this
// is a min- rather than max-priority queue.

type internalItem[T any] struct {
	value    T
	priority float64
}

type internalQueue[T any] []*internalItem[T]

func (q internalQueue[T]) Len() int { return len(q) }

func (q internalQueue[T]) Less(i, j int) bool {
	// Unlike the docs example, this is a min priority queue so we use less than.
	return q[i].priority < q[j].priority
}

func (q internalQueue[T]) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *internalQueue[T]) Push(x any) {
	item := x.(*internalItem[T])
	*q = append(*q, item)
}

func (q *internalQueue[T]) Pop() any {
	old := *q
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // don't stop the GC from reclaiming the item eventually
	*q = old[0 : n-1]
	return item
}
