package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func main() {
	http.HandleFunc("/", helloWorld)
	fmt.Println("Server started and listening on localhost")
	log.Fatal(http.ListenAndServe(":9003", nil))
}
