package main

type Node struct {
	data interface{}
	next *Node
}

type LinkedList struct {
	head *Node
	size int
}

func (l *LinkedList) Add(data interface{}) {
	newNode := &Node{data, l.head}
	l.head = newNode
	l.size++
}

func (l *LinkedList) Remove() interface{} {
	if l.head == nil {
		return nil
	}
	data := l.head.data
	l.head = l.head.next
	l.size--
	return data
}
