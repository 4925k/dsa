package main

import (
	"fmt"
)

type node struct {
	data int
	prev *node
	next *node
}

type list struct {
	head  *node
	tail  *node
	size  int
	count int
}

func main() {
	list := list{size: 5}
	list.insertLast(5)
	list.insertLast(6)
	list.insertLast(7)
	list.insertLast(8)
	list.insertLast(9)
	list.insertFirst(10)
	list.insertMiddle(111)
	list.deleteFirst()
	list.deleteLast()
	list.deleteMiddle()
	list.display()
}

func (list *list) insertLast(data int) {
	if list.isEmpty() {
		node := node{data: data}
		list.head, list.tail = &node, &node
		node.next, node.prev = &node, &node
		list.count++
	} else {
		node := node{data: data, next: list.head, prev: list.tail}
		list.tail.next, list.head.prev = &node, &node
		list.tail = &node
		list.count++
	}
}

func (list *list) insertFirst(data int) {
	if list.isEmpty() {
		node := node{data: data}
		list.head, list.tail = &node, &node
		node.next, node.prev = &node, &node
		list.count++
	} else {
		node := node{data: data, next: list.head, prev: list.tail}
		list.head.prev, list.tail.next = &node, &node
		list.head = &node
		list.count++
	}
}

func (list *list) insertMiddle(data int) {
	if list.isEmpty() {
		node := node{data: data}
		list.head, list.tail = &node, &node
		node.next, node.prev = &node, &node
		list.count++
	} else {
		centerNode := list.getNode((list.size + 1) / 2)
		node := node{data: data, next: centerNode.next, prev: centerNode}
		centerNode.next.prev = &node
		centerNode.next = &node
		list.count++
	}
}

func (list *list) deleteFirst() {
	list.head = list.head.next //new head
	list.tail.next = list.head //last elements new next
	list.head.prev = list.tail //current head new prev
	list.count--
}

func (list *list) deleteLast() {
	list.tail = list.tail.prev
	list.tail.next = list.head
	list.head.prev = list.tail
	list.count--

}

func (list *list) deleteMiddle() {
	beforeCenterNode := list.getNode(list.size / 2)
	beforeCenterNode.next.next.prev = beforeCenterNode
	beforeCenterNode.next = beforeCenterNode.next.next
	list.count--
}

func (list *list) display() {
	node := list.head
	var count int
	for node != nil {
		fmt.Println(node)
		node = node.next
		count++
		if list.count == count {
			return
		}
	}
}

func (list *list) getNode(n int) *node {
	node := list.head
	for i := 0; i < n-1; i++ {
		node = node.next
	}
	return node
}

func (list *list) isEmpty() bool {
	return list.head == nil && list.tail == nil
}
