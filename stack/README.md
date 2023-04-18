# Stack in Go

This project contains an implementation of a stack in Go. The stack is implemented using a `Stack` struct with a slice of interface{} values.

## Usage

To use the stack in your Go project, simply copy the `stack.go` file from this project into your project and import it as needed.

```go
package main

import (
	"fmt"
)

func main() {
	stack := Stack{}
	stack.Push(1)
	stack.Push(2)
	fmt.Println(stack.Pop()) // Output: 2
	fmt.Println(stack.Pop()) // Output: 1
	fmt.Println(stack.Pop()) // Output: nil
}
```