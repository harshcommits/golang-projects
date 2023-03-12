package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter text")
	/*
		input, _ := reader.ReadString('\n') //here, _ is used as substitue for error variable
		fmt.Println("You entered: ", input)
	*/

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("You entered: ", input)
	}

}
