package heaps

import "github.com/njayp/heaps/internal/generic"

// NewGenericHeap creates a new generic heap.
// `less` determines the heap type:
// - For a min-heap, use `less = func(a, b T) bool { return a < b }`.
// - For a max-heap, use `less = func(a, b T) bool { return a > b }`.
// Optionally, you can prefill the heap with initial data.
func NewGenericHeap[T any](less func(a, b T) bool, initialData ...T) *Heap[T] {
	return NewHeap[T](generic.NewHeapImpl(less, initialData...))
}
