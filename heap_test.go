package heaps

import (
	"testing"
)

// expect is a helper function that pops an element from the heap and checks if it matches the expected value.
func expect(t *testing.T, fn func() int, expected int) {
	t.Helper()

	val := fn()
	if val != expected {
		t.Errorf("expected %v, got %v", expected, val)
	}
}

func TestMinHeap(t *testing.T) {
	h := NewMinHeap(3, 1, 4)

	expect(t, h.Size, 3)
	expect(t, h.Peek, 1)
	expect(t, h.Pop, 1)
	expect(t, h.Peek, 3)

	h.Push(2)
	h.Push(5)

	expect(t, h.Size, 4)
	expect(t, h.Pop, 2)
	expect(t, h.Peek, 3)
	expect(t, h.Pop, 3)
	expect(t, h.Peek, 4)
	expect(t, h.Pop, 4)
	expect(t, h.Peek, 5)
	expect(t, h.Pop, 5)
	expect(t, h.Size, 0)
}

func TestMaxHeap(t *testing.T) {
	h := NewMaxHeap(0, 3, 1, 2)

	expect(t, h.Pop, 3)
	expect(t, h.Pop, 2)
	expect(t, h.Pop, 1)
	expect(t, h.Pop, 0)
	expect(t, h.Size, 0)
}
