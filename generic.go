package heaps

import "github.com/njayp/heaps/internal/generic"

// NewGenericHeap initializes a new generic heap.
// `less` determines the heap type:
// - For a min-heap, use `less = func(a, b T) bool { return a < b }`.
// - For a max-heap, use `less = func(a, b T) bool { return a > b }`.
func NewGenericHeap[T any](less func(a, b T) bool, initialData ...T) *Heap[T] {
	h := NewHeap[T](generic.NewHeapImpl(less, initialData...))
	if len(initialData) > 0 {
		h.Init()
	}
	return h
}
