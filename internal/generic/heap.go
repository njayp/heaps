package generic

import "container/heap"

// HeapImpl is a generic heap structure.
type HeapImpl[T any] struct {
	data []T
	less func(a, b T) bool // Comparison function to determine heap order
}

// NewHeapImpl impl the heap interface
func NewHeapImpl[T any](less func(a, b T) bool, initialData ...T) *HeapImpl[T] {
	h := &HeapImpl[T]{
		less: less,
	}
	if len(initialData) > 0 {
		h.data = initialData
		heap.Init(h) // Initialize heap with prefilled data
	}

	return h
}

func (h HeapImpl[T]) Len() int           { return len(h.data) }
func (h HeapImpl[T]) Less(i, j int) bool { return h.less(h.data[i], h.data[j]) }
func (h HeapImpl[T]) Swap(i, j int)      { h.data[i], h.data[j] = h.data[j], h.data[i] }

func (h *HeapImpl[T]) Push(x interface{}) {
	h.data = append(h.data, x.(T))
}

func (h *HeapImpl[T]) Pop() interface{} {
	old := h.data
	n := len(old)
	x := old[n-1]
	h.data = old[:n-1]
	return x
}