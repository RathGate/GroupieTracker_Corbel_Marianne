package main

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/packages/api"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
)

type Data struct {
	DataType     string
	PerfectMatch api.Item
	ResultArr    []api.Item
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		perfectMatch, err := api.MakeEntryRequest(r.FormValue("name"))
		allResults, err2 := api.SearchByName(true, r.FormValue("name"), perfectMatch.ID)
		if err != nil || err2 != nil {
			log.Fatal(err, err2)
		}
		tmpl := template.Must(template.ParseFiles("templates/index.html", "templates/_card_item.html"))
		tmpl.Execute(w, Data{DataType: "search", PerfectMatch: perfectMatch, ResultArr: allResults})
		return
	}
	result, err := api.MakeFullRequest()
	if err != nil {
		fmt.Println(err)
	}
	tmpl := template.Must(template.ParseFiles("templates/index.html", "templates/_card_item.html"))
	tmpl.Execute(w, Data{DataType: "all", ResultArr: result})
}

func main() {
	static := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", static))

	// Handles routing:
	http.HandleFunc("/", indexHandler)

	// Launches the server:
	preferredPort := ":8080"
	fmt.Printf("Starting server at port %v\n", preferredPort)
	if err := http.ListenAndServe(preferredPort, nil); err != nil {
		log.Fatal(err)
	}
}

func GenerateFallback() error {
	result, err := api.MakeFullRequest()
	if err != nil {
		return err
	}
	file, err := json.MarshalIndent(result, "", "   ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("assets/data/fallback.json", file, 0644)
	return err
}
