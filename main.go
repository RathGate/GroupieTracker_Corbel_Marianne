package main

import (
	"fmt"
	"groupie-tracker/packages/api"
	"log"
	"net/http"
	"text/template"
)

type Data struct {
	DataType  string
	ResultArr []api.Item
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// if r.Method == "POST" {
	// 	if result, err := api.MakeEntryRequest(r.FormValue("name")); err != nil {
	// 		fmt.Println(err)
	// 	} else if result.Item.Name != "" {
	// 		tmpl := template.Must(template.ParseFiles("templates/index.html", "templates/_card_item.html"))
	// 		tmpl.Execute(w, Data{DataType: "arr", ResultArr: []api.Item{result.Item}})
	// 		return
	// 	}
	// }
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
