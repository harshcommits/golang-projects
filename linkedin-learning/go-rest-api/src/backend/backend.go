package backend

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	Port   string
	Router *mux.Router
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from backend")
}

func (a *App) Initialize() {

	a.Router = mux.NewRouter()
	a.initializeRoutes()

}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/", helloWorld)
}

func (a *App) Run() {
	fmt.Println("Server started and listening on port", a.Port)
	log.Fatal(http.ListenAndServe(a.Port, a.Router))
}
