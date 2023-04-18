# Linked List in Go

This project contains an implementation of a singly linked list in Go. The linked list is implemented using a `Node` struct and a `LinkedList` struct. Each `Node` contains an interface{} value and a pointer to the next node in the list. The `LinkedList` struct contains a pointer to the head node and the size of the list.

## Usage

To use the linked list in your Go project, simply copy the `linkedlist.go` file from this project into your project and import it as needed.

```go
package main

import (
"fmt"
)

func main() {
    list := LinkedList{}
    list.Add(1)
    list.Add(2)
    fmt.Println(list.Remove()) // Output: 2
    fmt.Println(list.Remove()) // Output: 1
    fmt.Println(list.Remove()) // Output: nil
}

```