package main

import (
	"fmt"
)

func main() {

	fmt.Println("Hello there!")

	//multi-line declaration
	var data int32
	data = 32
	fmt.Printf("Multi-line declaration: %d\n", data)

	//single-line declaration
	other := 34
	fmt.Printf("Single line declaration: %d", other)

}
