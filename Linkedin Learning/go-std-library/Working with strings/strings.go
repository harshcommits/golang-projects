package main

import (
	"fmt"
	"strings"
)

func main() {

	s := "Testing strings"

	//print length
	fmt.Println(len(s))

	//iterate over each letter
	for _, ch := range s {
		fmt.Print(string(ch), ",")
	}
	fmt.Print("\n")

	//using operators < > == !=
	fmt.Println("dog" < "cat")   //false; since c before d
	fmt.Println("dog" < "horse") //true
	fmt.Println("dog" == "Dog")  //false since not same

	//comparing strings
	result := strings.Compare("dog", "cat")
	fmt.Println(result) //returns 1; hence false
	result2 := strings.Compare("dog", "dog")
	fmt.Println(result2) //returns 0; hence true

	//using EqualFold
	fmt.Println(strings.EqualFold("dog", "Dog")) //returns true

	//convert to lower, upper
	fmt.Println(strings.ToLower(s))
	fmt.Println(strings.ToUpper(s))
	fmt.Println(strings.Title(s)) //first letter uppercase

}
