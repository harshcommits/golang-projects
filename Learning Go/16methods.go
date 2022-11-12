package main

import (
	"fmt"
)

type Dog struct {
	Breed  string
	Weight int
	Sound  string
}

func main() {

	poodle := Dog{"Poodle", 10, "Woof!"}
	fmt.Printf("%+v\n", poodle)

	poodle.Speak()
	poodle.Sound = "Arf!"
	poodle.Speak()
	poodle.SpeakThreeTimes()

}

func (d Dog) Speak() { //d is receiver, where d is an instance of Dog struct and can be called in the method itself
	fmt.Println(d.Sound)
}

func (d Dog) SpeakThreeTimes() {
	/*
		d.Sound = fmt.Sprintf("%v %v %v", d.Sound, d.Sound, d.Sound)
		fmt.Println(d.Sound)
	*/

	for i := 0; i < 3; i++ {
		fmt.Println(d.Sound)
	}
}
