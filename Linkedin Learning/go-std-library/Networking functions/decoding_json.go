package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name       string   `json:"fullname"`
	Address    string   `json:"addr"`
	Age        int      `json:"myage"`
	FaveColors []string `json:"myfavecolors"`
}

func decodeExample() {

	data := []byte(`
		{
			"fullname" : "John Doe",
			"addr" : "987 Main Str",
			"myage" : 45,
			"myfavecolors" : ["Purple", "White", "Gold"]
		}
	`)

	//var to decode struct
	var p person

	//check if json is valid
	valid := json.Valid(data)
	if valid {
		json.Unmarshal(data, &p)
		fmt.Printf("%#v\n", p)
	}

	//decode to a map structure
	var m map[string]interface{}

	//unmarshal in a map
	json.Unmarshal(data, &m)
	fmt.Printf("%#v\n", m)

	for k, v := range m {
		fmt.Printf("key (%v), value (%T : %v)\n", k, v, v)
	}

}

func main() {
	decodeExample()
}
