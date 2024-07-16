package main

import (
	"database/sql"
	"fmt"
	"log"

	// "net/http"

	_ "github.com/mattn/go-sqlite3"
)

type Product struct {
	id        int
	name      string
	inventory int
	price     int
}

// func helloWorld(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello World")
// }

func main() {
	// http.HandleFunc("/", helloWorld)
	// fmt.Println("Server started and listening on localhost")
	// log.Fatal(http.ListenAndServe(":9003", nil))

	db, err := sql.Open("sqlite3", "./practiceit.db")
	if err != nil {
		log.Fatal(err.Error())
	}

	rows, err := db.Query("SELECT id, name, inventory, price FROM products")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var p Product

		rows.Scan(&p.id, &p.name, &p.inventory, &p.price)
		fmt.Println("Product, ", p.id, p.name, p.inventory, p.price)
	}

}
