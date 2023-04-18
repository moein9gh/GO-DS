# Bloom Filter in Go

This project contains an implementation of a Bloom filter in Go. The Bloom filter is implemented using a `BloomFilter` struct with a bitset and a hash function.

## Usage

To use the Bloom filter in your Go project, simply copy the `bloomfilter.go` file from this project into your project and import it as needed.

```go
package main

import (
	"fmt"
)

func main() {
	bf := NewBloomFilter(100)
	bf.Add([]byte("hello"))
	bf.Add([]byte("world"))
	fmt.Println(bf.Contains([]byte("hello"))) // Output: true
	fmt.Println(bf.Contains([]byte("world"))) // Output: true
	fmt.Println(bf.Contains([]byte("test")))  // Output: false
}
```