package heaps

import (
	"testing"
)

// expectPop is a helper function that pops an element from the heap and checks if it matches the expected value.
func expectPop(t *testing.T, h *Heap[int], expected int) {
	t.Helper()
	val := h.Pop()
	if val != expected {
		t.Errorf("expected %v, got %v", expected, val)
	}
}

// TestHeap_Push tests the Push method of the heap.
func TestHeap_Push(t *testing.T) {
	h := NewMinHeap()
	h.Push(3)
	h.Push(1)
	h.Push(2)

	expectPop(t, h, 1)
	expectPop(t, h, 2)
	expectPop(t, h, 3)
}

// TestHeap_Pop tests the Pop method of the heap.
func TestHeap_Pop(t *testing.T) {
	h := NewMinHeap(3, 1, 2)

	expectPop(t, h, 1)
	expectPop(t, h, 2)
	expectPop(t, h, 3)
}

// TestHeap_Init tests the initialization of the heap.
func TestHeap_Init(t *testing.T) {
	h := NewMinHeap(3, 1, 2)

	expectPop(t, h, 1)
	expectPop(t, h, 2)
	expectPop(t, h, 3)
}

// TestMaxHeap_Push tests the Push method of the max heap.
func TestMaxHeap_Push(t *testing.T) {
	h := NewMaxHeap()
	h.Push(1)
	h.Push(3)
	h.Push(2)

	expectPop(t, h, 3)
	expectPop(t, h, 2)
	expectPop(t, h, 1)
}

type Item struct {
	Priority int
}

// TestHeap_Fix tests the Fix method of the heap with custom priority items.
func TestHeap_Fix(t *testing.T) {
	d := &Item{Priority: 2}
	h := NewGenericHeap(func(a, b *Item) bool { return a.Priority < b.Priority },
		d,
		&Item{Priority: 1},
		&Item{Priority: 3},
	)
	d.Priority = 0 // directly modify the underlying data
	h.Fix(1)       // fix the heap

	got := h.Pop()
	if got.Priority != 0 {
		t.Errorf("expected 0, got %v", got.Priority)
	}
	got = h.Pop()
	if got.Priority != 1 {
		t.Errorf("expected 2, got %v", got.Priority)
	}
	got = h.Pop()
	if got.Priority != 3 {
		t.Errorf("expected 3, got %v", got.Priority)
	}
}
