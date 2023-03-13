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
var lastRequest LastRequest
var lastPage string

type LastRequest struct {
	Page       string
	AllResults []api.Item
}

func (request *LastRequest) updateRequest(page string, results []api.Item) {
	lastRequest.Page = page
	lastRequest.AllResults = results
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	lastPage = "index"
	data := Data{PageName: lastPage}

	tmpl := template.Must(template.ParseFiles("templates/base.html", "templates/views/index.html"))
	tmpl.Execute(w, data)
}
func itemHandler(w http.ResponseWriter, r *http.Request) {
	lastPage = "item"

	id := mux.Vars(r)["id"]
	item, err := api.MakeEntryRequest(id)
	if err != nil {
		fmt.Println(err)
		return
	}

	data := Data{PageName: lastPage, PerfectMatch: item}

	tmpl := template.Must(template.ParseFiles("templates/base.html", "templates/views/item.html", "templates/components/_single_page.html"))
	tmpl.Execute(w, data)
}
func categoriesHandler(w http.ResponseWriter, r *http.Request) {
	lastPage = "item"
	data := Data{PageName: lastPage}

	if r.Method == "POST" {
		var temp []api.Item
		if categoryStr := r.FormValue("category-id"); categoryStr != "" {
			categoryID, err := strconv.Atoi(categoryStr)
			if err != nil || categoryID < 0 || categoryID > 4 {
				return
			}

			temp, err = api.MakeCategoryRequest(damn[categoryID])
			if err != nil {
				println(err)
				return
			}

			tmpl, _ := template.New("").ParseFiles("templates/components/_card_item.html")

			err = tmpl.ExecuteTemplate(w, "card", temp)

			if err != nil {
				panic(err)
			}
			lastRequest.updateRequest("idk", temp)
			return
		}
	}
	temp, err := api.MakeCategoryRequest("creatures")
	if err != nil {
		println(err)
		return
	}

	if err != nil {
		log.Fatal(err)
		return
	}
	data.ResultArr = temp
	tmpl := template.Must(template.ParseFiles("templates/base.html", "templates/views/categories.html", "templates/components/_card_item.html"))
	tmpl.Execute(w, data)

	fmt.Println(len(lastRequest.AllResults))
	lastRequest.updateRequest("idk", temp)
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
	fmt.Println(len(lastRequest.AllResults))
}
