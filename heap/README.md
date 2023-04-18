# Heap in Go

This project contains an implementation of a binary max heap in Go. The heap is implemented using a `Heap` struct with a slice of `int` values.

## Usage

To use the heap in your Go project, simply copy the `heap.go` file from this project into your project and import it as needed.

```go
package main

import (
	"fmt"
)

func main() {
	heap := Heap{}
	heap.Push(5)
	heap.Push(10)
	heap.Push(7)
	fmt.Println(heap.Pop()) // Output: 10
	fmt.Println(heap.Pop()) // Output: 7
	fmt.Println(heap.Pop()) // Output: 5
	fmt.Println(heap.Pop()) // Output: -1
}
```