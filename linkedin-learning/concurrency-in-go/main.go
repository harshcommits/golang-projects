package main

import (
	"fmt"
	"sync"
	"time"
)

var greetings = []string{"Hello", "Ciao", "Namaste"}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)     // update waitgroup counter
	go Hello(&wg) // run goroutine with waitgroup
	wg.Wait()     // wait for the waitgroup counter to update
	goodbye()

	// server.Start()

	ch := make(chan string) //waits for 5s for line 23 to execute, and the "greeter completed" message
	// ch := make(chan string, 1) // channel buffer set to one value; "greeter completed" without having to wait for line 23 execution
	// go greet(ch)
	go greet2(ch)

	// time.Sleep(5 * time.Second)
	// fmt.Println("Main ready:")
	// greeting := <-ch
	// time.Sleep(2 * time.Second)
	// fmt.Println("Greeting received")
	// fmt.Println(greeting)

	time.Sleep(1 * time.Second)
	fmt.Println("Main 2 ready")

	// iterating through channels with fix for avoiding empty loops and check if the channel is closed
	// for {
	// 	greeting, ok := <-ch // check with the ok value if the channel is closed
	// 	if !ok {
	// 		return
	// 	}

	// 	time.Sleep(500 * time.Millisecond)
	// 	fmt.Println("Greeting received: ", greeting)
	// }

	// better way to iterate through channels
	for greeting := range ch {
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Greeting received: ", greeting)
	}

}

func greet2(ch chan<- string) {
	fmt.Println("Greeter ready")
	for _, g := range greetings {
		ch <- g
	}
	close(ch) // close channel to avoid goroutine locks
	fmt.Println("Greeter completed")
}

// func greet(ch chan<- string) { // chan<- is a unidirectional channel (best practices); you can only send values in it
// 	fmt.Printf("Greeter ready and waiting...")
// 	ch <- "Hello World"
// 	fmt.Println("Greeter completed")
// }

func Hello(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Hello World")
}

func goodbye() {
	fmt.Println("goodbye")
}
