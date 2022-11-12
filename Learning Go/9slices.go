package main

import (
	"fmt"
	"sort"
)

func main() {
	var colors = []string{"Red", "Blue", "Green"}
	fmt.Println(colors)
	colors = append(colors, "Purple")
	fmt.Println(colors)

	//colors = append(colors[1:len(colors)])
	//fmt.Println(colors)
	colors = append(colors[:len(colors)-1])
	fmt.Println(colors)

	numbers := make([]int, 6) //creates a slice of 6 values; unless declared, each value is 0
	numbers[0] = 123
	numbers[1] = 23
	numbers[2] = 34
	numbers[3] = 57

	fmt.Println(numbers)

	numbers = append(numbers, 332)
	fmt.Println(numbers)
	sort.Ints(numbers) //sort function for int slices
	fmt.Println("sorted: ", numbers)

}
