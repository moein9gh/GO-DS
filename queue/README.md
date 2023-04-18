# Queue in Go

This project contains an implementation of a queue in Go. The queue is implemented using a `Queue` struct with a slice of interface{} values.

## Usage

To use the queue in your Go project, simply copy the `queue.go` file from this project into your project and import it as needed.

```go
package main

import (
	"fmt"
)

func main() {
	queue := Queue{}
	queue.Enqueue(1)
	queue.Enqueue(2)
	fmt.Println(queue.Dequeue()) // Output: 1
	fmt.Println(queue.Dequeue()) // Output: 2
	fmt.Println(queue.Dequeue()) // Output: nil
}
```