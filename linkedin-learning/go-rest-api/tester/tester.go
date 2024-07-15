package main

import (
	"fmt"

	"github.com/golang-projects/linkedin-learning/go-rest-api/funcvmet"
	"github.com/golang-projects/linkedin-learning/go-rest-api/practice"
)

func main() {
	practice.Test()

	//done via receiver methods
	d := funcvmet.Dimensions{Length: 1, Width: 2, Height: 4} // can be done without keys; shows unkeyed composites warning tho
	fmt.Println(d.Area())
	fmt.Println(d.Volume())

	// done via functions
	area, vol := funcvmet.Values(1, 2, 4)
	fmt.Println(area, vol)
}
