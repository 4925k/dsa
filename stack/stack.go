package main

import "fmt"

//stack stores stack details
//making a slice removes the limitation of size.
//so adding size to work with the given limitation on a stack size.
type stack struct {
	top   int
	size  int
	stack []string
}

func main() {
	s := createStack(2)
	s.push("1")
	fmt.Println(s.stack)
	s.push("2")
	fmt.Println(s.stack)
	s.push("3")
	fmt.Println(s.stack)
	fmt.Println("Using peek:", s.peek())
	fmt.Println("Using pop:", s.pop())

}

//createStack returns a new stack
func createStack(size int) stack {
	return stack{
		top:   -1,
		size:  size,
		stack: make([]string, size),
	}
}

//push adds element to the top of the stack
func (s *stack) push(item string) {
	if s.isFull() {
		fmt.Println("The stack is full")
		return
	}

	s.top++
	s.stack[s.top] = item
	fmt.Printf("Added %s to stack\n", item)
}

//pop removes the top element from the stack
func (s *stack) pop() string {
	if s.isEmpty() {
		fmt.Println("The stack is empty")
		return ""
	}
	element := s.stack[s.top]
	s.stack = s.stack[:(s.top)]
	fmt.Printf("Removed %s from the stack\n", element)
	return element
}

//isEmpty returns true if stack is empty else false.
func (s *stack) isEmpty() bool {
	return s.top == -1
}

//isFull returns true if the stack if full else false.
func (s *stack) isFull() bool {
	return s.top == (s.size - 1)
}

func (s *stack) peek() string {
	return s.stack[s.top]
}
