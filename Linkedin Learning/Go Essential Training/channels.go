package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	// ch <- 353 //send; this throws error

	go func() {
		//send value to channel
		ch <- 353
	}()

	val := <-ch //receive
	fmt.Printf("Value from channel: %d\n", val)

	//send multiple values to channel
	const count = 3
	go func() {
		for i := 0; i < count; i++ {
			fmt.Printf("Sending %d\n", i)
			ch <- i
			time.Sleep(time.Second)
		}
	}()

	for i := 0; i < count; i++ {
		val := <-ch
		fmt.Printf("received %d\n", val)
	}
}
