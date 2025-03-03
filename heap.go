package heaps

import (
	"container/heap"
)

// NewHeap initializes a new generic heap.
// `less` determines the heap type:
// - For a min-heap, use `less = func(a, b T) bool { return a < b }`.
// - For a max-heap, use `less = func(a, b T) bool { return a > b }`.
func NewHeap[T any](less func(a, b T) bool, initialData ...T) Heap[T] {
	h := Heap[T]{
		h: &genericHeapImpl[T]{
			data: initialData,
			less: less,
		},
	}

	if len(initialData) > 0 {
		heap.Init(h.h)
	}

	return h
}

// Heap is a wrapper for a generic heap implementation.
type Heap[T any] struct {
	h *genericHeapImpl[T]
}

// Push pushes the element x onto the heap. The complexity is O(log n) where n = h.Len().
func (h Heap[T]) Push(t T) {
	heap.Push(h.h, t)
}

// Pop removes and returns the minimum element (according to Less) from the heap. The complexity is O(log n) where n = h.Len(). Panics if the heap is empty.
func (h Heap[T]) Pop() T {
	return heap.Pop(h.h).(T)
}

// Peek returns the minimum element (according to Less) from the heap without removing it. The complexity is O(1). Panics if the heap is empty.
func (h Heap[T]) Peek() T {
	return h.h.data[0]
}

// Size returns the number of elements in the heap.
func (h Heap[T]) Size() int {
	return h.h.Len()
}
