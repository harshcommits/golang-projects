package main

import (
	"fmt"

	ds "github.com/go-ds/ds"
)

func main() {
	fmt.Println("This is a test")

	stack1 := ds.Stack{}
	stack1.Push(100)
	stack1.Push(20)
	stack1.Push(3)
	fmt.Println(stack1)
}
