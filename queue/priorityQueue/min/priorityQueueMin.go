package main

import "fmt"

type heap struct {
	queue []int
}

func main() {
	heap := heap{}
	data := []int{13, 24, 5, 16, 34, 8, 3, 1}
	for _, v := range data {
		heap.insert(v)
		fmt.Println(heap)
	}

	for i := 0; i < 3; i++ {
		heap.extract()
		fmt.Println(heap)
	}

	fmt.Println("Using peek to get min value", heap.peek())

	fmt.Println("extracting min:", heap.extract(), heap.queue)

}

func (h *heap) insert(element int) {
	h.queue = append(h.queue, element)
	h.heapifyUp(len(h.queue) - 1)
}

func (h *heap) extract() int {
	if len(h.queue) == 0 {
		fmt.Println("The queue is empty")
		return -1
	}
	element := h.queue[0]
	lastIndex := len(h.queue) - 1
	h.queue[0] = h.queue[lastIndex]
	h.queue = h.queue[:lastIndex]
	h.heapifyDown(0)
	return element
}

func (h *heap) peek() int {
	if len(h.queue) == 0 {
		fmt.Println("The queue is empty")
		return -1
	}
	return h.queue[0]
}

func (h *heap) heapifyDown(index int) {
	left := index*2 + 1
	right := index*2 + 2
	var childToCompare int
	for left < len(h.queue) {
		if left == (len(h.queue) - 1) {
			childToCompare = left
		} else if h.queue[left] < h.queue[right] {
			childToCompare = left
		} else {
			childToCompare = right
		}
		if h.queue[childToCompare] < h.queue[index] {
			h.queue[childToCompare], h.queue[index] = h.queue[index], h.queue[childToCompare]
			index = childToCompare
			left = index*2 + 1
			right = index*2 + 2
		} else {
			return
		}
	}
}

func (h *heap) heapifyUp(index int) {
	if index == 0 {
		return
	}
	parent := (index - 1) / 2
	if h.queue[index] < h.queue[parent] {
		h.queue[index], h.queue[parent] = h.queue[parent], h.queue[index]
		index = parent
		h.heapifyUp(index)
	}

}
