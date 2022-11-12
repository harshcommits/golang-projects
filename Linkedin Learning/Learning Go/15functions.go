package main

import (
	"fmt"
)

func main() {
	doSomething()
	fmt.Println(addValues(1, 2))
	multiSum, multiCount := addAllValues(1, 23, 3)
	fmt.Println(multiSum)
	fmt.Println("No. of values:", multiCount)
}

func doSomething() {
	fmt.Println("Doing something")
}

func addValues(value1 int, value2 int) int {
	return value1 + value2
}

//multiple values using single parameter
func addAllValues(values ...int) (int, int) {
	total := 0
	for _, v := range values {
		total += v
	}
	return total, len(values)
}
