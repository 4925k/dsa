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
	head *node
	tail *node
}

func main() {
	list := list{}
	list.insertLast(5)
	list.insertLast(6)
	list.insertLast(7)
	list.insertLast(8)
	list.insertLast(9)
	list.display()
	list.insertFirst(10)
	list.display()
	list.insertMiddle(111)
	list.display()
	list.deleteFirst()
	list.display()
	list.deleteLast()
	list.display()
	list.deleteMiddle()
	list.display()
}

//add if empty condition
func (list *list) insertLast(data int) {
	node := node{data: data}
	if list.head == nil {
		list.head, list.tail = &node, &node
	} else {
		list.tail.next = &node
		node.prev = list.tail
		list.tail = &node
	}
}

func (list *list) insertFirst(data int) {
	node := node{data: data, next: list.head}
	list.head.prev = &node
	list.head = &node
}

func (list *list) insertMiddle(data int) {
	centerNode := list.getNode((list.count() + 1) / 2)
	node := node{data: data, next: centerNode.next, prev: centerNode}
	centerNode.next.prev = &node
	centerNode.next = &node
}

func (list *list) deleteFirst() {
	list.head = list.head.next
	list.head.prev = nil
}

func (list *list) deleteLast() {
	node := list.tail.prev
	list.tail = node
	node.next = nil
}

func (list *list) deleteMiddle() {
	beforeCenterNode := list.getNode(list.count() / 2)
	beforeCenterNode.next.next.prev = beforeCenterNode
	beforeCenterNode.next = beforeCenterNode.next.next
}

func (list *list) display() {
	node := list.head
	for node != nil {
		fmt.Println(node.data)
		node = node.next
	}
}

func (list *list) getNode(n int) *node {
	node := list.head
	for i := 0; i < n-1; i++ {
		node = node.next
	}
	return node
}

func (list *list) count() int {
	node := list.head
	var count int
	for node != nil {
		count++
		node = node.next
	}
	return count
}
