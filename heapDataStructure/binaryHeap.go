package main

import "fmt"

type heap struct {
	array []int
}

func main() {
	x := &heap{}
	y := []int{10, 20, 30, 5, 7, 9, 11, 13, 15, 17}
	for _, v := range y {
		x.insert(v)
		fmt.Println(x)
	}
	for i := 0; i < 5; i++ {
		x.extract()
		fmt.Println(x)
	}
}
func (h *heap) insert(key int) {
	h.array = append(h.array, key)
	h.heapifyUp(len(h.array) - 1)
}

func (h *heap) heapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

func (h *heap) heapifyDown(index int) {
	lastIndex := len(h.array) - 1
	l, r := left(index), right(index)
	childToCompare := 0
	for l <= lastIndex {
		if l == lastIndex {
			childToCompare = l
		} else if h.array[l] > h.array[r] {
			childToCompare = l
		} else {
			childToCompare = r
		}
		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}
}
func (h *heap) extract() int {
	if len(h.array) == 0 {
		fmt.Println("the array length is 0")
		return -1
	}
	extracted := h.array[0]
	length := len(h.array) - 1
	h.array[0] = h.array[length]
	h.array = h.array[:length]

	h.heapifyDown(0)
	return extracted
}

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return (i * 2) + 1
}

func right(i int) int {
	return (i * 2) + 2
}
func (h *heap) swap(i, j int) {
	h.array[i], h.array[j] = h.array[j], h.array[i]
}
