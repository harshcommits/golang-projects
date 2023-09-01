package main

import (
	"fmt"

	ds "github.com/go-ds/ds"
)

func main() {

	// stack1 := ds.Stack{}
	// stack1.Push(100)
	// stack1.Push(20)
	// stack1.Push(3)
	// fmt.Println(stack1)

	// fmt.Println("break")
	// stack1.Pop()
	// fmt.Println(stack1)

	// queue1 := ds.Queue{}
	// queue1.Enqueue(100)
	// queue1.Enqueue(2)
	// queue1.Enqueue(30)
	// fmt.Println(queue1)

	// queue1.Dequeue()
	// fmt.Println(queue1)

	ll1 := ds.LinkedList{}
	node1 := &ds.Node{Data: 49}
	node2 := &ds.Node{Data: 35}
	node3 := &ds.Node{Data: 14}

	ll1.Prepend(node1)
	ll1.Prepend(node2)
	ll1.Prepend(node3)
	fmt.Println(ll1)
	ll1.PrintListData()

	ll1.DeleteWithValue(14) // deleting the head
	ll1.PrintListData()
	ll1.DeleteWithValue(100)

}
