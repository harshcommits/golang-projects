package ds

import (
	"fmt"
)

type Node struct {
	Data int
	next *Node
}

type LinkedList struct {
	Head   *Node
	length int
}

func (l *LinkedList) Prepend(n *Node) {
	second := l.Head
	l.Head = n
	l.Head.next = second
	l.length++
}

func (l LinkedList) PrintListData() {
	toPrint := l.Head
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.Data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Printf("\n")
}

func (l *LinkedList) DeleteWithValue(value int) error {

	// case for empty list
	if l.length == 0 {
		return fmt.Errorf("empty list")
	}

	// if value matches the head of the list
	if l.Head.Data == value {
		l.Head = l.Head.next
		l.length--
		return nil
	}

	previousToDelete := l.Head
	for previousToDelete.next.Data != value {

		if previousToDelete.next.next == nil {
			return fmt.Errorf("reached the end: value not in the list") // exits out without returning anything
		}
		previousToDelete = previousToDelete.next
	}
	previousToDelete.next = previousToDelete.next.next
	l.length--

	return nil
}
