package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name       string   `json:"fullname"` //shows value in backticks
	Address    string   `json:"addr"`
	Age        int      `json:"-"`                   //age won't be outputted
	FaveColors []string `json:"favcolors,omitempty"` //doesn't show null in output if empty
}

func encodeExample() {

	people := []person{
		{"Jane Doe", "123 Anywhere Street", 35, nil},
		{"John Doe", "456 Everywhere Blvd", 31, []string{"Purple", "Yellow"}},
	}

	result, err := json.MarshalIndent(people, "", "\t") //sets up indentation for json
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", result)

}

func main() {
	encodeExample()
}
