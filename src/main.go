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

// This program generates fallback files on server launch.
// These files are backup .json files that contain all processed
// data for endpoints /all and master_mode/all and are used on /search route.
// !Feel free to disable this feature if you encounter any problem with the search result,
// !but please note that the request time might be increased.
const USEFALLBACK = true

func main() {
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

	// Generates fallback files:
	if USEFALLBACK {
		api.GenerateFallback(true, true)
	}

	// Launches the server:
	preferredPort := ":8080"
	fmt.Printf("Starting server at port %v\n", preferredPort)
	if err := http.ListenAndServe(preferredPort, r); err != nil {
		log.Fatal(err)
	}
}
