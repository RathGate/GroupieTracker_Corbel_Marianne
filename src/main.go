package main

import (
	"fmt"
	"groupie-tracker/packages/api"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Data struct {
	PageName     string
	DataType     string
	PerfectMatch api.Item
	ResultArr    []api.Item
}

func main() {
	r := mux.NewRouter()

	// This will serve files under http://localhost:8000/assets/<filename>
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	// Handles routing:
	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/item/{id}", itemHandler)
	r.HandleFunc("/search", searchHandler)
	r.HandleFunc("/categories", categoriesHandler)
	r.NotFoundHandler = http.HandlerFunc(notFoundHandler)
	// Launches the server:
	preferredPort := ":8080"
	fmt.Printf("Starting server at port %v\n", preferredPort)
	if err := http.ListenAndServe(preferredPort, r); err != nil {
		log.Fatal(err)
	}
}
