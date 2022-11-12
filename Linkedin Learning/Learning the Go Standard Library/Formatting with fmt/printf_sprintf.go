package main

import (
	"fmt"
)

type circle struct {
	radius int
	border int
}

func main() {

	x := 20
	f := 123.45

	fmt.Printf("%d\n", x)
	fmt.Printf("%x\n", f) //%x gives hex value

	//get boolean value
	fmt.Printf("%t\n", x > 10)

	//floats
	fmt.Printf("%f\n", f)
	fmt.Printf("%e\n", f) // %e gives exponential values

	//print explicit values
	fmt.Printf("%[2]d %[1]d\n", 21, 42) //prints 42 and then 21

	//print same value repeatedly
	fmt.Printf("%d %#[1]o %#[1]x\n", 21) //prints 21 as int, octal and hex

	//print custom types
	c := circle{
		radius: 20,
		border: 5,
	}
	fmt.Printf("%+v\n", c) //using + adds field name in output
	fmt.Printf("%T\n", c)  //using %T gives type

	//sprintf returns the string; doesn't print it
	//can be put for internal use, without printing anything
	s := fmt.Sprintf("%[2]d %[1]d\n", 52, 40)
	fmt.Println(s)

}
