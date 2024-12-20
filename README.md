# Heaps
## Some Boilerplate Code for Heaps
The goal is to build upon `container/heap` for easier usage. Specifically: 

- remove the need for boilerplate code
- reduce the api footprint

To accomplish this, this library implements a generic array that satisfies the `heap` interface. This array is then wrapped with `container/heap` functionality.

## Instillation

```sh
go get github.com/njayp/heaps
```

## Usage
Create a heap of custom structs

```go
    // define custom struct
	type item struct {
		value int
	}

    // a function for comparing structs
    less := func(a, b item) bool {
		return a.value < b.value
	}

    // init heap with data
	heap := heaps.NewHeap(less, 
    item{value: 1},
    item{value: 2},
    )

    // use heap
	min := heap.Pop() // item{value: 1}
```

Or use a preset heap

```go
    // init heap with data
    heap := heaps.NewMaxHeap(0, 1, 2)

    // use heap
	max := heap.Pop() // 2
```

## Additional Functionality
Using a generic array allows us to add `Peek` functionality, as the root of the heap will always be at index zero.