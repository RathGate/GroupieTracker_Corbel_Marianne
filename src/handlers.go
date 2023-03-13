package main

import (
	"fmt"
	"groupie-tracker/packages/api"
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/gorilla/mux"
)

var damn = []string{"creatures", "monsters", "materials", "equipment", "treasure"}
var lastRequest []api.Item

func indexHandler(w http.ResponseWriter, r *http.Request) {
	data := Data{PageName: "index"}
	tmpl := template.Must(template.ParseFiles("templates/base.html", "templates/views/index.html"))
	tmpl.Execute(w, data)
}
func itemHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	item, err := api.MakeEntryRequest(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := Data{PageName: "item", PerfectMatch: item}
	tmpl := template.Must(template.ParseFiles("templates/base.html", "templates/views/item.html", "templates/components/_single_page.html"))
	tmpl.Execute(w, data)
}
func categoriesHandler(w http.ResponseWriter, r *http.Request) {
	data := Data{PageName: "categories"}

	if r.Method == "POST" {
		var temp []api.Item
		if categoryStr := r.FormValue("category-id"); categoryStr != "" {
			categoryID, err := strconv.Atoi(categoryStr)
			if err != nil || categoryID < 0 || categoryID > 4 {
				return
			}
			lastRequest, err = api.MakeCategoryRequest(damn[categoryID])

			if err != nil {
				println(err)
				return
			}
			if len(lastRequest) > 24 {
				temp = lastRequest[:24]
			} else {
				temp = lastRequest
			}

			tmpl, _ := template.New("").ParseFiles("templates/components/_card_item.html")
			err = tmpl.ExecuteTemplate(w, "card", temp)
			if err != nil {
				panic(err)
			}
			return
		}
	}
	lastRequest, err := api.MakeCategoryRequest("creatures")
	if err != nil {
		log.Fatal(err)
		return
	}
	if len(lastRequest) > 24 {
		data.ResultArr = lastRequest[:24]
	} else {
		data.ResultArr = lastRequest
	}
	tmpl := template.Must(template.ParseFiles("templates/base.html", "templates/views/categories.html", "templates/components/_card_item.html"))
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
