# Binary Search Tree in Go

This project contains an implementation of a binary search tree in Go. The binary search tree is implemented using a `BinarySearchTree` struct with a `Node` struct.

## Usage

To use the binary search tree in your Go project, simply copy the `bst.go` file from this project into your project and import it as needed.

```go
package main

import (
"fmt"
)

func main() {
    bst := BinarySearchTree{}
    bst.Insert(5)
    bst.Insert(10)
    bst.Insert(7)
    fmt.Println(bst.Search(7)) // Output: true
    fmt.Println(bst.Search(12)) // Output: false
}
```