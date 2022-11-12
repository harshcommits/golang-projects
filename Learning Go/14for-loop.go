package main

import (
	"fmt"
)

func main() {
	colors := []string{"Red", "Green", "Blue"}
	fmt.Println(colors)

	/*
		//standard for loop
		for i := 0; i < len(colors); i++ {
			fmt.Println(colors[i])
		}

		//for loop using range
		for i := range colors {
			fmt.Println(colors[i])
		}

		//for loop mapping as an index and corresponding value; _ is placeholder for index
		for _, color := range colors {
			fmt.Println(color)
		}
	*/

	value := 1
	for value < 10 {
		fmt.Println("Value: ", value)
		value++
	}

	//using goto labels to jump to other sections
	sum := 1
	for sum < 1000 {
		sum += sum
		fmt.Println("Sum: ", sum)
		if sum < 200 {
			goto theEnd
		}

	}

theEnd:
	fmt.Println("End of program")

}
