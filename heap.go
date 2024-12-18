package heaps

import (
	"container/heap"
)

func NewHeap[T any](h heap.Interface) *Heap[T] {
	return &Heap[T]{h: h}
}

type Heap[T any] struct {
	h heap.Interface
}

// Push pushes the element x onto the heap. The complexity is O(log n) where n = h.Len().
func (h *Heap[T]) Push(t T) {
	heap.Push(h.h, t)
}

// Pop removes and returns the minimum element (according to Less) from the heap. The complexity is O(log n) where n = h.Len(). Pop is equivalent to [Remove](h, 0).
func (h *Heap[T]) Pop() T {
	return heap.Pop(h.h).(T)
}

// Remove removes and returns the element at index i from the heap. The complexity is O(log n) where n = h.Len().
func (h *Heap[T]) Remove(i int) {
	heap.Remove(h.h, i)
}

// Fix re-establishes the heap ordering after the element at index i has changed its value. Changing the value of the element at index i and then calling Fix is equivalent to, but less expensive than, calling [Remove](h, i) followed by a Push of the new value. The complexity is O(log n) where n = h.Len().
func (h *Heap[T]) Fix(i int) {
	heap.Fix(h.h, i)
}

// Init establishes the heap invariants required by the other routines in this package. Init is idempotent with respect to the heap invariants and may be called whenever the heap invariants may have been invalidated. The complexity is O(n) where n = h.Len().
func (h *Heap[T]) Init() {
	heap.Init(h.h)
}
