package main

import (
	"fmt"
	"groupie-tracker/packages/api"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	// !To uncomment before real launch
	// Generates fallback files:
	api.GenerateFallback(true, true)
	rand.Seed(time.Now().UnixNano())

	// Handlers router creation and static files:
	r := mux.NewRouter()
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
