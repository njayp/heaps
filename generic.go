package heaps

// genericHeapImpl implements the heap interface.
type genericHeapImpl[T any] struct {
	data []T
	less func(a, b T) bool // Comparison function to determine heap order
}

func (h genericHeapImpl[T]) Len() int           { return len(h.data) }
func (h genericHeapImpl[T]) Less(i, j int) bool { return h.less(h.data[i], h.data[j]) }
func (h genericHeapImpl[T]) Swap(i, j int)      { h.data[i], h.data[j] = h.data[j], h.data[i] }

func (h *genericHeapImpl[T]) Push(x interface{}) {
	h.data = append(h.data, x.(T))
}

func (h *genericHeapImpl[T]) Pop() interface{} {
	n := h.Len()
	x := h.data[n-1]
	h.data = h.data[:n-1]
	return x
}
