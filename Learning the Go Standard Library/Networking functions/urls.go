package main

import (
	"fmt"
	"net/url"
)

func main() {

	s := "https://www.example.com:8000/user?username=joemarini"

	//use the parse function, to get different parts of url
	result, _ := url.Parse(s)
	fmt.Println(result.Scheme)   //prints https
	fmt.Println(result.Host)     //print website
	fmt.Println(result.Path)     //prints /user
	fmt.Println(result.Port())   //prints 8000
	fmt.Println(result.RawQuery) //prints username=joemarini

	//extract query in a values struct
	vals := result.Query()
	fmt.Println(vals["username"])

	//create URL from components
	newurl := &url.URL{
		Scheme:   "https",
		Host:     "www.example.com",
		Path:     "/args",
		RawQuery: "x=1&y=2",
	}
	s = newurl.String()
	fmt.Println(s)

	//modify URL components
	newurl.Host = "joemarini.com"
	s = newurl.String()
	fmt.Println(s)

	//create new values struct and encode it as query string
	newvals := url.Values{}
	newvals.Add("x", "100")
	newvals.Add("z", "somestr")

	newurl.RawQuery = newvals.Encode() //generates the query string; x=100&z=somestr
	s = newurl.String()
	fmt.Println(s)

}
