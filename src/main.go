package main

import (
	"fmt"
	"groupie-tracker/packages/api"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Data struct {
	PageName     string
	DataType     string
	PerfectMatch api.Item
	ResultArr    []api.Item
	Regions      []string
	Categories   []string
	Mastermode   bool
}

var REGIONS_NAMES = []string{"Akkala", "Central Hyrule", "Eldin", "Faron", "Gerudo", "Hebra", "Lanayru", "Necluda"}

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
	GenerateMMFallback()
	// Launches the server:
	preferredPort := ":8080"
	fmt.Printf("Starting server at port %v\n", preferredPort)
	if err := http.ListenAndServe(preferredPort, r); err != nil {
		log.Fatal(err)
	}

}

func applyFilters(filters Filters) (result []api.Item) {
	var allitems []api.Item
	var err error
	if filters.MasterMode {
		allitems, _ = api.UseFallBack(true)
	} else {
		allitems, err = api.UseFallBack(false)
		if err != nil {
			fmt.Println(err)
		}
	}

	for _, item := range allitems {
		if strings.Contains(item.Name, filters.Name) && isInRegion(filters.Regions, item) && stringInSlice(item.Category, filters.Category) {
			result = append(result, item)
		}
	}
	return result
}
