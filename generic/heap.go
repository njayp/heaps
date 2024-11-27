package generic

import "container/heap"

// GenericHeap is a generic heap structure.
type GenericHeap[T any] struct {
	data []T
	less func(a, b T) bool // Comparison function to determine heap order
}

// NewGenericHeap creates a new generic heap.
// `less` determines the heap type:
// - For a min-heap, use `less = func(a, b T) bool { return a < b }`.
// - For a max-heap, use `less = func(a, b T) bool { return a > b }`.
// Optionally, you can prefill the heap with initial data.
func NewGenericHeap[T any](less func(a, b T) bool, initialData ...T) *GenericHeap[T] {
	h := &GenericHeap[T]{
		data: []T{},
		less: less,
	}
	if len(initialData) > 0 {
		h.data = initialData
		heap.Init(h) // Initialize heap with prefilled data
	}
	return h
}

// Implement heap.Interface methods for GenericHeap[T].

func (h GenericHeap[T]) Len() int           { return len(h.data) }
func (h GenericHeap[T]) Less(i, j int) bool { return h.less(h.data[i], h.data[j]) }
func (h GenericHeap[T]) Swap(i, j int)      { h.data[i], h.data[j] = h.data[j], h.data[i] }

func (h *GenericHeap[T]) Push(x interface{}) {
	h.data = append(h.data, x.(T))
}

func (h *GenericHeap[T]) Pop() interface{} {
	old := h.data
	n := len(old)
	x := old[n-1]
	h.data = old[:n-1]
	return x
}

// Add an additional method to get the heap top.
func (h *GenericHeap[T]) Top() T {
	if len(h.data) == 0 {
		var zero T
		return zero
	}
	return h.data[0]
}
