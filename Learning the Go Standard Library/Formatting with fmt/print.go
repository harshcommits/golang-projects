package main

import "fmt"

func main() {

	//no new line created at end
	fmt.Print("Welcome to Go")

	fmt.Println("This string ends with a new line") // \n not required

	const answer = 42
	fmt.Println("The number is:", 42) //adds space between string and 42 automatically

	//print string with multiple values
	const a, b, c = 5, 5, 10
	fmt.Println("Add", a, "and", b, "gets", c)

	//print a slice
	items := []int{10, 23, 30, 40, 50}
	length, err := fmt.Println(items)
	fmt.Println(length, err)

}
