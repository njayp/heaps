package heaps

func NewMinHeap(initialData ...int) *Heap[int] {
	return NewHeap(func(a, b int) bool { return a < b }, initialData...)
}

func NewMaxHeap(initialData ...int) *Heap[int] {
	return NewHeap(func(a, b int) bool { return a > b }, initialData...)
}
