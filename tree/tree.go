package main

type TreeNode struct {
	children map[rune]*TreeNode
	isEnd    bool
}

type Tree struct {
	root *TreeNode
}

func (t *Tree) Insert(word string) {
	if t.root == nil {
		t.root = &TreeNode{children: make(map[rune]*TreeNode)}
	}
	current := t.root
	for _, c := range word {
		if _, ok := current.children[c]; !ok {
			current.children[c] = &TreeNode{children: make(map[rune]*TreeNode)}
		}
		current = current.children[c]
	}
	current.isEnd = true
}

func (t *Tree) Search(word string) bool {
	current := t.root
	for _, c := range word {
		if _, ok := current.children[c]; !ok {
			return false
		}
		current = current.children[c]
	}
	return current.isEnd
}
