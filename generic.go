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
	old := h.data
	n := len(old)
	x := old[n-1]
	h.data = old[:n-1]
	return x
}
