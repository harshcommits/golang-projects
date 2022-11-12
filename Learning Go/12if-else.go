package main

import (
	"fmt"
)

func main() {

	theAnswer := 42
	var result string

	if theAnswer < 0 {
		result = "less than zero"
	} else if theAnswer == 0 {
		result = "Equal to zero"
	} else {
		result = "Greater than zero"
	}
	fmt.Println(result)

	if x := -42; x < 0 { //adding condition and statement in same line
		result = "Less than zero"
	} else if x == 0 {
		result = "Equal to zero"
	} else {
		result = "Greater than zero"
	}
	fmt.Println(result)

}
