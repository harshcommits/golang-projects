package main

import (
	"fmt"
	"strings"
)

func main() {

	var sb strings.Builder

	//write something
	sb.WriteString("This is a string 1\n")
	sb.WriteString("This is a string 2\n")
	sb.WriteString("This is a string 3\n")
	sb.WriteString("This is a string 4\n")

	//output string
	fmt.Println(sb.String())

	//
	fmt.Println("Capacity: ", sb.Cap())   //shows storage capacity of string; here it is 96
	sb.Grow(1024)                         //add capacity by 1024
	fmt.Println("New Capacity", sb.Cap()) //here it shows 1216

	for i := 0; i <= 10; i++ {
		fmt.Fprintf(&sb, "String %d - ", i)
	}

	//print string builder size
	fmt.Println("the builder size is", sb.Len()) //size is 198

	//reset string builder space
	sb.Reset()
	fmt.Println("After reset, size is: ", sb.Cap()) //size is 0 here
}
