package main

import (
	"fmt"
	"groupie-tracker/packages/api"
	"log"
	"net/http"
)

type Data struct {
	PageName     string
	DataType     string
	PerfectMatch api.Item
	ResultArr    []api.Item
}

func main() {
	static := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", static))

	// Handles routing:
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/search", searchHandler)
	http.HandleFunc("/categories", categoriesHandler)

	// Launches the server:
	preferredPort := ":8080"
	fmt.Printf("Starting server at port %v\n", preferredPort)
	if err := http.ListenAndServe(preferredPort, nil); err != nil {
		log.Fatal(err)
	}
}
