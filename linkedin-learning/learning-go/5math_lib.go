package main

import (
	"fmt"
	"math"
)

func main() {

	i1, i2, i3 := 12, 45, 48
	intSum := i1 + i2 + i3
	fmt.Println("The sum of ints: ", intSum)

	f1, f2, f3 := 21.2, 34.3, 33.1
	floatSum := f1 + f2 + f3
	fmt.Println("Sum of floats: ", floatSum)

	circleRadius := 15.5
	circumference := 2 * math.Pi * circleRadius
	fmt.Printf("Circumference: %.2f\n", circumference) //.2f denotes value to 2 decimal points

}
