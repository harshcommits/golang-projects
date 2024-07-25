package main

import (
	"fmt"
	"time"
)

func main() {

	go sayHello("Hello")
	// sayHello("Hello")
	sayHello("World")

}

func sayHello(message string) {
	for i := 0; i < 5; i++ {
		fmt.Println(message)
		time.Sleep(time.Second)
	}
}
