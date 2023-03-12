package main

import (
	"fmt"
)

func main() {
	anInt := 42
	var p = &anInt //p is the address for &anInt, where the value is stored
	fmt.Println("Value of p: ", *p)

	value1 := 42.13
	pointer1 := &value1
	fmt.Println("Value 1: ", *pointer1)
}
