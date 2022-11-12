package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func getRequestTest() {

	const httpbin = "https://httpbin.org/get"

	//perform get operation
	resp, err := http.Get(httpbin)
	if err != nil {
		return
	}

	//close the response by defer
	defer resp.Body.Close()

	//print status from response
	fmt.Println("Status: ", resp.Status)
	fmt.Println("Status code: ", resp.StatusCode)
	fmt.Println("Protocol: ", resp.Proto)
	fmt.Println("Content Length: ", resp.ContentLength)

	//create stringbuilder for content
	var sb strings.Builder
	content, _ := ioutil.ReadAll(resp.Body)
	bytecount, _ := sb.Write(content)

	//print bytecount
	fmt.Println(bytecount, sb.String())

}

func main() {
	getRequestTest()
}
