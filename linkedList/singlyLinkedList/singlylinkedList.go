package main

import "fmt"

type node struct {
	data int
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
	list.insertFirst(10)
	list.insertMiddle(111)
	list.deleteFirst()
	list.deleteLast()
	list.deleteMiddle()
	list.display()
}

func (list *list) insertLast(data int) {
	node := node{data: data}
	if list.head == nil {
		list.head, list.tail = &node, &node
	} else {
		list.tail.next = &node
		list.tail = &node
	}
}

func (list *list) insertFirst(data int) {
	node := node{data: data, next: list.head}
	list.head = &node
}

func (list *list) insertMiddle(data int) {

	centerNode := list.getNode((list.count() + 1) / 2)
	node := node{data: data, next: centerNode.next}
	centerNode.next = &node
}

func (list *list) deleteFirst() {
	list.head = list.head.next
}

func (list *list) deleteLast() {
	node := list.head
	for node.next.next != nil {
		node = node.next
	}
	node.next = nil
}

func (list *list) deleteMiddle() {
	beforeCenterNode := list.getNode(list.count() / 2)
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
