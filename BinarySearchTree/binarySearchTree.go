package main

type Node struct {
	data  int
	left  *Node
	right *Node
}

type BinarySearchTree struct {
	root *Node
}

func (t *BinarySearchTree) Insert(data int) {
	if t.root == nil {
		t.root = &Node{data: data}
		return
	}
	current := t.root
	for {
		if data < current.data {
			if current.left == nil {
				current.left = &Node{data: data}
				return
			}
			current = current.left
		} else if data > current.data {
			if current.right == nil {
				current.right = &Node{data: data}
				return
			}
			current = current.right
		} else {
			return
		}
	}
}

func (t *BinarySearchTree) Search(data int) bool {
	current := t.root
	for current != nil {
		if current.data == data {
			return true
		} else if current.data > data {
			current = current.left
		} else {
			current = current.right
		}
	}
	return false
}
