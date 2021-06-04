package main

import "fmt"

type queue struct {
	front int
	rear  int
	size  int
	queue []string
}

func main() {
	q := createQueue(4)
	fmt.Println(q)

	q.enqueue("1")
	fmt.Println(q)
	q.enqueue("2")
	fmt.Println(q)
	q.enqueue("3")
	fmt.Println(q)
	q.enqueue("4")
	fmt.Println(q)
	q.dequeue()
	fmt.Println(q)
	q.dequeue()
	fmt.Println(q)
	q.enqueue("5")
	fmt.Println(q)
	q.enqueue("6")
	fmt.Println(q)
	q.enqueue("7")
	fmt.Println(q)
	q.enqueue("8")
	fmt.Println(q)
	q.dequeue()
	fmt.Println(q)
	q.enqueue("9")
	fmt.Println(q)

}

func createQueue(size int) queue {
	return queue{
		front: -1,
		rear:  -1,
		size:  size,
		queue: make([]string, size),
	}
}

func (q *queue) enqueue(element string) {
	if q.isFull() {
		fmt.Println("The queue is full")
		return
	}
	if q.front == -1 { //if element is the first in the queue
		q.front = 0
	}
	q.rear++
	if q.rear == q.size {
		q.rear = 0
	}
	q.queue[q.rear] = element
}

func (q *queue) dequeue() string {
	if q.isEmpty() {
		fmt.Println("The queue is empty")
		return ""
	}
	element := q.queue[q.front]
	q.queue[q.front] = ""
	q.front++
	if q.front == q.rear {
		q.front = -1
		q.rear = -1
	}
	return element
}

func (q *queue) isFull() bool {
	if q.front == 0 && q.rear == q.size-1 {
		return true
	}
	if q.front == (q.rear + 1) {
		return true
	}
	return false
}

func (q *queue) isEmpty() bool {
	return q.front == -1
}
