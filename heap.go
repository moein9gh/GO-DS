package main

type Heap struct {
	items []int
}

func (h *Heap) ParentIndex(index int) int {
	return (index - 1) / 2
}

func (h *Heap) LeftChildIndex(index int) int {
	return 2*index + 1
}

func (h *Heap) RightChildIndex(index int) int {
	return 2*index + 2
}

func (h *Heap) Swap(index1, index2 int) {
	h.items[index1], h.items[index2] = h.items[index2], h.items[index1]
}

func (h *Heap) Peek() int {
	return h.items[0]
}

func (h *Heap) Push(item int) {
	h.items = append(h.items, item)
	currentIndex := len(h.items) - 1
	parentIndex := h.ParentIndex(currentIndex)
	for currentIndex != 0 && h.items[currentIndex] > h.items[parentIndex] {
		h.Swap(currentIndex, parentIndex)
		currentIndex = parentIndex
		parentIndex = h.ParentIndex(currentIndex)
	}
}

func (h *Heap) Pop() int {
	if len(h.items) == 0 {
		return -1
	}
	lastIndex := len(h.items) - 1
	item := h.items[0]
	h.items[0] = h.items[lastIndex]
	h.items = h.items[:lastIndex]
	currentIndex := 0
	leftChildIndex := h.LeftChildIndex(currentIndex)
	rightChildIndex := h.RightChildIndex(currentIndex)
	for leftChildIndex < lastIndex {
		maxIndex := leftChildIndex
		if rightChildIndex < lastIndex && h.items[rightChildIndex] > h.items[maxIndex] {
			maxIndex = rightChildIndex
		}
		if h.items[currentIndex] > h.items[maxIndex] {
			break
		}
		h.Swap(currentIndex, maxIndex)
		currentIndex = maxIndex
		leftChildIndex = h.LeftChildIndex(currentIndex)
		rightChildIndex = h.RightChildIndex(currentIndex)
	}
	return item
}
