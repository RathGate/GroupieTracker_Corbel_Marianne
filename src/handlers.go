package main

import (
	"fmt"
	"groupie-tracker/packages/api"
	"log"
	"net/http"
	"text/template"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	data := Data{PageName: "index"}
	tmpl := template.Must(template.ParseFiles("templates/base.html", "templates/views/index.html"))
	tmpl.Execute(w, data)
}
func categoriesHandler(w http.ResponseWriter, r *http.Request) {
	data := Data{PageName: "categories"}
	tmpl := template.Must(template.ParseFiles("templates/base.html", "templates/views/categories.html"))
	tmpl.Execute(w, data)
}
func searchHandler(w http.ResponseWriter, r *http.Request) {
	data := Data{PageName: "search"}
	if r.Method == "POST" {
		perfectMatch, err := api.MakeEntryRequest(r.FormValue("name"))
		allResults, err2 := api.SearchByName(true, r.FormValue("name"), perfectMatch.ID)
		if err != nil || err2 != nil {
			log.Fatal(err, err2)
		}
		data.PerfectMatch = perfectMatch
		data.ResultArr = allResults
		tmpl := template.Must(template.ParseFiles("templates/base.html", "templates/views/search.html"))
		tmpl.Execute(w, data)
		return
	}
	result, err := api.UseFallBack()
	if err != nil {
		fmt.Println(err)
	}
	data.ResultArr = result
	tmpl := template.Must(template.ParseFiles("templates/base.html", "templates/views/search.html"))
	tmpl.Execute(w, data)
}
