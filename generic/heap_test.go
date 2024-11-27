package generic

import (
	"container/heap"
	"testing"
)

// TestGenericHeap tests the GenericHeap implementation.
func TestGenericHeap(t *testing.T) {
	// Test 1: Min-Heap for integers
	minHeap := NewGenericHeap(func(a, b int) bool { return a < b })
	heap.Init(minHeap)

	heap.Push(minHeap, 10)
	heap.Push(minHeap, 5)
	heap.Push(minHeap, 20)

	if got := heap.Pop(minHeap).(int); got != 5 {
		t.Errorf("Min-Heap Pop: got %d, want %d", got, 5)
	}
	if got := heap.Pop(minHeap).(int); got != 10 {
		t.Errorf("Min-Heap Pop: got %d, want %d", got, 10)
	}
	if got := heap.Pop(minHeap).(int); got != 20 {
		t.Errorf("Min-Heap Pop: got %d, want %d", got, 20)
	}

	// Test 2: Max-Heap for integers
	maxHeap := NewGenericHeap(func(a, b int) bool { return a > b })
	heap.Init(maxHeap)

	heap.Push(maxHeap, 10)
	heap.Push(maxHeap, 5)
	heap.Push(maxHeap, 20)

	if got := heap.Pop(maxHeap).(int); got != 20 {
		t.Errorf("Max-Heap Pop: got %d, want %d", got, 20)
	}
	if got := heap.Pop(maxHeap).(int); got != 10 {
		t.Errorf("Max-Heap Pop: got %d, want %d", got, 10)
	}
	if got := heap.Pop(maxHeap).(int); got != 5 {
		t.Errorf("Max-Heap Pop: got %d, want %d", got, 5)
	}

	// Test 3: Custom struct with a min-heap
	type Person struct {
		Name string
		Age  int
	}

	// init an existing heap. it's faster
	personHeap := NewGenericHeap(func(a, b Person) bool { return a.Age > b.Age },
		Person{Name: "Alice", Age: 30},
		Person{Name: "Bob", Age: 25},
		Person{Name: "Charlie", Age: 35},
	)

	if got := heap.Pop(personHeap).(Person); got.Name != "Bob" {
		t.Errorf("Person Min-Heap Pop: got %s, want %s", got.Name, "Bob")
	}
	if got := heap.Pop(personHeap).(Person); got.Name != "Alice" {
		t.Errorf("Person Min-Heap Pop: got %s, want %s", got.Name, "Alice")
	}
	if got := heap.Pop(personHeap).(Person); got.Name != "Charlie" {
		t.Errorf("Person Min-Heap Pop: got %s, want %s", got.Name, "Charlie")
	}
}
