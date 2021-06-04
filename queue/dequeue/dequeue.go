package main

import "fmt"

type dequeue struct {
	front int
	rear  int
	size  int
	queue []int
}

func main() {
	q := createDequeue(5)
	fmt.Println(q)

	q.insertFront(5)
	fmt.Println(q)
	q.insertFront(6)
	fmt.Println(q)
	q.insertFront(7)
	fmt.Println(q)
	q.insertRear(8)
	fmt.Println(q)
	q.insertRear(9)
	fmt.Println(q)
	q.insertRear(8)
	q.removeFront()
	fmt.Println(q)
	q.removeRear()
	fmt.Println(q)
	q.removeRear()
	fmt.Println(q)
	q.removeRear()
	fmt.Println(q)
	q.removeRear()
	fmt.Println(q)
	q.removeRear()
	fmt.Println(q)
}

func (d *dequeue) insertFront(element int) {
	if d.isFull() {
		fmt.Println("The queue is full")
		return
	}
	if d.front == -1 {
		d.front++
	} else if d.front == 0 {
		d.front = len(d.queue) - 1
	} else {
		d.front--
	}
	d.queue[d.front] = element
}

func (d *dequeue) insertRear(element int) {
	if d.isFull() {
		fmt.Println("The queue is full.")
		return
	}
	if d.front == -1 {
		d.front++
	} else if d.rear == len(d.queue)-1 {
		d.rear = 0
	} else {
		d.rear++
	}
	d.queue[d.rear] = element

}

func (d *dequeue) removeFront() {
	if d.isEmpty() {
		fmt.Println("The queue is empty.")
		return
	}
	current := d.front
	if d.front == d.rear {
		d.front = -1
		d.rear = 0
	} else if d.front == len(d.queue)-1 {
		d.front = 0
	} else {
		d.front++
	}
	d.queue[current] = 0
}

func (d *dequeue) removeRear() {
	if d.isEmpty() {
		fmt.Println("The queue is empty")
		return
	}
	current := d.rear
	if d.front == d.rear {
		d.front = -1
		d.rear = 0
	} else if d.rear == 0 {
		d.rear = len(d.queue) - 1
	} else {
		d.rear--
	}
	d.queue[current] = 0
}
func createDequeue(size int) dequeue {
	return dequeue{
		front: -1,
		rear:  0,
		queue: make([]int, size),
	}
}

func (d *dequeue) isFull() bool {
	return ((d.front == 0 && d.rear == d.size-1) || d.front == d.rear+1)
}

func (d *dequeue) isEmpty() bool {
	return d.front == -1
}
