package main

import (
	"fmt"
)

type Dog struct {
	Breed  string //uppercase keys; are accessible from outside the package
	Weight int
	size   string //only accessible inside the package
}

func main() {

	poodle := Dog{"Poodle", 10, "small"}
	fmt.Println(poodle)
	fmt.Printf("%+v\n", poodle) //shows data as key-value pairs, as per struct definition
	fmt.Printf("%#v\n", poodle) //shows type of struct with key:values; eg. main.Dog{Breed:"Poodle", Weight:10}
	fmt.Printf("Breed: %v\nWeight: %v\n", poodle.Breed, poodle.Weight)
	fmt.Printf("Size: %v\n", poodle.size)

}
