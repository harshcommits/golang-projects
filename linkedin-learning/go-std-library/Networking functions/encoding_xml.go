package main

import (
	"encoding/xml"
	"fmt"
)

type person struct {
	XMLName    xml.Name `xml:"persondata"`
	FirstName  string   `xml:"firstname"`
	LastName   string   `xml:"lastname"`
	Address    string   `xml:"addr"`
	Age        int      `xml:"age,attr"` //will be shown as attribute not tag
	FaveColors []string `xml:"favecolors"`
}

func main() {

	people := []person{
		{FirstName: "Jane", LastName: "Doe", Address: "123 Avenue", Age: 33, FaveColors: []string{}},
		{FirstName: "John", LastName: "Doe", Address: "Everywhere Blvd", Age: 22, FaveColors: []string{"Blue", "Green"}},
	}

	//marshalindent for xml conversion
	result, err := xml.MarshalIndent(people, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s%s\n", xml.Header, result)

}
