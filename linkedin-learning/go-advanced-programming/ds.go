package main

import (
	"errors"
	"fmt"
)

// stack definition
type Stack[T any] struct { // using notation for generics
	elements []T
}

type Node[T any] struct {
	value      T
	next, prev *Node[T]
}

type DoublyLinkedList[T any] struct {
	size int
	head *Node[T]
	tail *Node[T]
}

// function to push values in the stack
func (s *Stack[T]) Push(element T) {
	s.elements = append(s.elements, element)
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *Stack[T]) Pop() (*T, error) {
	if s.IsEmpty() {
		return nil, errors.New("Cannot pop from empty stack")
	}

	top := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]

	return &top, nil
}

func main() {
	fmt.Println("This is the area for coding data structures")

	// stack := Stack[int]{elements: []int{1, 2, 4, 5, 6}}
	stack := Stack[int]{}
	elements := []int{1, 2, 4, 7}
	// stack.Push(3)

	for _, value := range elements {
		stack.Push(value)
	}

	fmt.Println(stack.elements)
}
