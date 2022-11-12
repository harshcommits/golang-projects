package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func postRequestTest() {

	const httpbin = "https://httpbin.org/post"

	//post operation
	reqBody := strings.NewReader(`{
		"field1" : "This is field 1",
		"field2" : 250
	}
	`)

	resp, err := http.Post(httpbin, "application/json", reqBody)
	if err != nil {
		return
	}

	content, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	fmt.Printf("%s\n", content)

	fmt.Println("Using postForm")

	//postform
	data := url.Values{}
	data.Add("field1", "Field added via Values")
	data.Add("field2", "300")
	resp, err = http.PostForm(httpbin, data)
	content, _ = ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	fmt.Printf("%s\n", content)

}

func main() {
	postRequestTest()
}
