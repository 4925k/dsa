package main

import "fmt"

type queue struct {
	front int
	rear  int
	size  int
	queue []string
}

func createQueue(size int) queue {
	return queue{
		front: -1,
		rear:  -1,
		size:  size,
		queue: make([]string, size),
	}
}
func main() {
	queue := createQueue(3)

	queue.enqueue("1")
	fmt.Println(queue)

	queue.enqueue("2")
	fmt.Println(queue)

	queue.dequeue()
	fmt.Println(queue)

	fmt.Println("Using peek: ", queue.peek())
}

//enqueue adds the element to the queue
func (q *queue) enqueue(element string) {
	if q.isFull() {
		if !(q.front == 0) {
			fmt.Println("Empty the queue first")
			return
		}
		fmt.Println("The queue is full.")
		return
	}
	q.front = 0
	q.rear++
	q.queue[q.rear] = element
}

//dequeue removes the first element from the queue
func (q *queue) dequeue() string {
	if q.isEmpty() {
		fmt.Println("The queue is empty.")
		return ""
	}
	element := q.queue[q.front]
	q.queue[q.front] = ""
	q.front++
	if q.front == q.size {
		q.front = -1
		q.rear = -1
	}
	return element
}

//isEmpty returns true is queue is empty and false if not
func (q *queue) isEmpty() bool {
	return q.front == -1
}

//isFull returns true is queue is full and false if not
func (q *queue) isFull() bool {
	return q.rear == (q.size - 1)
}

//peel returns the first element in the queue
func (q *queue) peek() string {
	return q.queue[q.front]
}
