package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter text: ")
	input, _ := reader.ReadString('\n') // _ used as placeholder(disposable) variable for error; \n works as delimiter
	fmt.Println("You entered: ", input)

	fmt.Println("Enter a number: ")
	numInput, _ := reader.ReadString('\n')
	aFloat, err := strconv.ParseFloat(strings.TrimSpace(numInput), 64) //adding err since error logging required here
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Value of number: ", aFloat)
	}

}
