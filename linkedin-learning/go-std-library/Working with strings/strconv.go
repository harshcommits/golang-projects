package main

import (
	"fmt"
	"strconv"
)

func main() {

	sampleint := 100
	samplestr := "250"

	newstr := string(sampleint)
	fmt.Printf("Result of using %T: %v\n", newstr, newstr) //prints d instead of 100

	s := strconv.Itoa(sampleint)
	fmt.Printf("%T: %v\n", s, s)

	//convert string to integer
	sampleint, err := strconv.Atoi(samplestr)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%T, %v\n", sampleint, sampleint)

	//other conversion
	b, _ := strconv.ParseBool("true")
	fmt.Println(b)
	f, _ := strconv.ParseFloat("3.14159", 64)
	fmt.Println(f)
	i, _ := strconv.ParseInt("-42", 10, 64)
	fmt.Println(i)

}
