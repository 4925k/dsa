package main

import "fmt"

type tree struct {
	queue []int
}

func main() {
	heap := tree{}
	data := []int{1, 2, 5, 16, 34, 8, 3}
	for _, v := range data {
		heap.insert(v)
		fmt.Println(heap)
	}

	for i := 0; i < 3; i++ {
		heap.remove(0)
		fmt.Println(heap)
	}

	fmt.Println("Using peek to get max value", heap.peek())

	fmt.Println("extracting max:", heap.extractMax(), heap.queue)

}
func (t *tree) insert(element int) {
	t.queue = append(t.queue, element)
	t.heapifyUp(len(t.queue) - 1)
}

func (t *tree) remove(index int) {
	if len(t.queue) == 0 {
		fmt.Println("The queue is empty.")
		return
	}
	lastIndex := len(t.queue) - 1
	t.queue[0] = t.queue[lastIndex]
	t.queue = t.queue[:lastIndex]
	t.heapifyDown(index)
}

func (t *tree) peek() int {
	return t.queue[0]
}

func (t *tree) heapifyDown(index int) {
	lastIndex := len(t.queue) - 1
	leftChild := index*2 + 1
	rightChild := index*2 + 2
	childToCompare := 0
	for leftChild < len(t.queue) {
		if leftChild == lastIndex {
			childToCompare = leftChild
		} else if t.queue[leftChild] > t.queue[rightChild] {
			childToCompare = leftChild
		} else {
			childToCompare = rightChild
		}
		if t.queue[index] < t.queue[childToCompare] {
			t.queue[childToCompare], t.queue[index] = t.queue[index], t.queue[childToCompare]
			index = childToCompare
			leftChild = index*2 + 1
			rightChild = index*2 + 2
		} else {
			return
		}
	}
}

func (t *tree) heapifyUp(index int) {
	if index == 0 {
		return
	}
	parent := (index - 1) / 2
	if t.queue[parent] < t.queue[index] {
		t.queue[parent], t.queue[index] = t.queue[index], t.queue[parent]
		index = parent
		t.heapifyUp(index)
	}
}

func (t *tree) extractMax() int {
	max := t.queue[0]
	t.remove(0)
	return max
}
