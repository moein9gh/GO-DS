package main

type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

type Trie struct {
	root *TrieNode
}

func (t *Trie) Insert(word string) {
	if t.root == nil {
		t.root = &TrieNode{children: make(map[rune]*TrieNode)}
	}
	current := t.root
	for _, c := range word {
		if _, ok := current.children[c]; !ok {
			current.children[c] = &TrieNode{children: make(map[rune]*TrieNode)}
		}
		current = current.children[c]
	}
	current.isEnd = true
}

func (t *Trie) Search(word string) bool {
	current := t.root
	for _, c := range word {
		if _, ok := current.children[c]; !ok {
			return false
		}
		current = current.children[c]
	}
	return current.isEnd
}
