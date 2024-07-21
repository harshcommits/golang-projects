package main

import (
	"github.com/golang-projects/linkedin-learning/go-rest-api/src/backend"
)

func main() {

	a := backend.App{}
	a.Port = ":9003"
	a.Initialize()
	a.Run()
}
