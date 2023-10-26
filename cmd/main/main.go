package main

import (
	"log"
	"net/http"

	"github.com/IvanGelo1/go-todo-crud/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterTaskRoutes(router)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":9010", router))
}
