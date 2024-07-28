package main

import (
	"log"

	"github.com/harshcommits/golang-projects/linkedin-learning/microservice-go/internal/database"
	"github.com/harshcommits/golang-projects/linkedin-learning/microservice-go/internal/server"
)

func main() {

	db, err := database.NewDatabaseClient()
	if err != nil {
		log.Fatalf("failed to initialize database client, %s", err)
	}

	srv := server.NewEchoServer(db)
	if err := srv.Start(); err != nil {
		log.Fatal(err.Error())
	}

}
