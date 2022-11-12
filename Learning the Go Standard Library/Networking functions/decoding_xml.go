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
	Age        int      `xml:"age,attr"`
	FaveColors []string `xml:"favecolors"`
}

func main() {

	//sample xml data
	xmldata := `
	<persondata age="29">
		<firstname>John</firstname>
		<lastname>Public</lastname>
		<addr>456 Everywhere Blvd</addr>
		<favecolors>Purple</favecolors>
		<favecolors>Yellow</favecolors>
		<favecolors>Green</favecolors>
	</persondata>
	`
	//variable for xml structure
	var p person

	//unmarshal the xml
	xml.Unmarshal([]byte(xmldata), &p)
	fmt.Printf("%#v\n", p)

}
