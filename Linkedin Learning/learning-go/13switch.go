package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	dow := rand.Intn(7) + 1 //sets ceiling to 7
	fmt.Println("Day", dow)

	var result string
	switch dow {
	case 1:
		result = "It's Sunday"
		// fallthrough		//can be used like break; will run next statement even if this is true
	case 2:
		result = "It's Monday"
	default:
		result = "Some other day then"
	}
	fmt.Println(result)

}
