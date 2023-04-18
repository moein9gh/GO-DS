# Trie in Go

This project contains an implementation of a trie in Go. The trie is implemented using a `TreeNode` struct and a `Tree` struct. Each `TreeNode` contains a map of `rune` values to child nodes and a boolean indicating whether it is the end of a word. The `Tree` struct contains a pointer to the root node.

## Usage

To use the trie in your Go project, simply copy the `trie.go` file from this project into your project and import it as needed.

```go
package main

import (
	"fmt"
)

func main() {
	trie := Tree{}
	trie.Insert("hello")
	trie.Insert("world")
	fmt.Println(trie.Search("hello")) // Output: true
	fmt.Println(trie.Search("world")) // Output: true
	fmt.Println(trie.Search("test"))  // Output: false
}
```