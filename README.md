# heaps
## some boilerplate code for heaps
The goal is to reduce the boilerplate code and api footprint of `container/heap` heaps. To accomplish this, a generic array that satisfies the `heap` interface is wrapped with  `container/heap` functionality.

## additional functionality
Using a generic array allows us to add Peek() functionality, as the root of the heap will always be at index zero.