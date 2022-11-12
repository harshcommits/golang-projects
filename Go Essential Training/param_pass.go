package main

import (
	"fmt"
)

func main() {
	values := []int{1, 2, 3, 4}
	doubleAt(values, 2)
	fmt.Println(values) //references are passed in cases of slices, so the values are changed even outside of function

	val := 10
	double(val)
	fmt.Printf("Value from outside the function: %d\n", val)

	doublePtr(&val)
	fmt.Printf("Value after executing the pointer function: %d", val)
}

func doubleAt(values []int, i int) {
	values[i] *= 2
}

func double(n int) {
	n *= 2
	fmt.Printf("Value from inside function: %d\n", n)
}

func doublePtr(n *int) {
	*n *= 2
	fmt.Printf("Value from inside the pointer function: %d\n", *n)
}
