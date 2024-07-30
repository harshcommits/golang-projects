package server

import (
	"fmt"
	"net/http"
)

func Start() {
	http.HandleFunc("/hello", hello)
	fmt.Println("Listening on port 8080")

	http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}
