package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n') //reads string until it encounters newline, or enter key input
	fmt.Println(s)

}
