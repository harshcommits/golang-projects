package main

import (
	"fmt"
)

const aConst int = 64

func main() {
	var aString string = "This is Golang on Linux"
	fmt.Println(aString)
	fmt.Printf("The variable type is %T\n", aString) //%T gives the type of variable; string in this case

	var anInt int = 42
	fmt.Println(anInt)

	//int by default has value zero
	var defaultInt int
	fmt.Println(defaultInt)

	var anotherString = "This is another string\n"
	fmt.Println(anotherString)
	fmt.Printf("The type is %T\n", anotherString) //determines type even if not specified

	//direct assignment; can be done only inside of function, outside requires var keyword
	myString := "This is a string without var"
	fmt.Println(myString)

	fmt.Printf("The type for const is %T", aConst)
}
