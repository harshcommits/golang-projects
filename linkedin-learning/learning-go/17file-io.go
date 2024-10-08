package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	content := "Hello from Go!"
	file, err := os.Create("./fromString.txt")
	checkError(err)

	length, err := io.WriteString(file, content)
	checkError(err)

	fmt.Printf("Wrote a file with %v characters\n", length)
	defer file.Close() //defer waits until all is done
	defer readFile("./fromString.txt")

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func readFile(fileName string) {
	data, err := os.ReadFile(fileName)
	checkError(err)
	fmt.Println("Text read from file: \n", string(data))

}
