package main

import (
	"fmt"
	"groupie-tracker/packages/api"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
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

	tmpl := template.Must(template.ParseFiles("templates/base.html", "templates/views/item.html", "templates/components/entry-item.html"))
	tmpl.Execute(w, data)
}
func testHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.RequestURI())
	reg := regexp.MustCompile(`\/test\?id=(?P<id>\d+)`)
	fmt.Println(reg.MatchString(r.URL.RequestURI()))
	//? if !reg.MatchString(r.URL.RequestURI()) {
	//! 	404
	//? }
	id := reg.FindStringSubmatch(r.URL.RequestURI())[1]

	item, err := api.MakeEntryRequest(id)
	if err != nil {
		fmt.Println(err)
		return
	}

	data := Data{PageName: lastPage, PerfectMatch: item}

	tmpl := template.Must(template.ParseFiles("templates/base.html", "templates/views/item.html", "templates/components/entry-item.html"))
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
			for i, v := range temp {
				fmt.Println(i, v.Name, v.ID)
			}
			tmpl, err := template.New("").ParseFiles("templates/components/card-item-container.html", "templates/components/card-item.html")
			fmt.Println(err)

			err = tmpl.ExecuteTemplate(w, "card-container", temp)

			if err != nil {
				fmt.Println(err)
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
	tmpl := template.Must(template.ParseFiles("templates/base.html", "templates/views/categories.html", "templates/components/card-item.html"))
	tmpl.Execute(w, data)

	fmt.Println(len(lastRequest.AllResults))
	lastRequest.updateRequest("idk", temp)
}
func searchHandler(w http.ResponseWriter, r *http.Request) {
	data := Data{PageName: "search", Regions: REGIONS_NAMES, Categories: damn}
	if r.Method == "POST" {
		r.ParseMultipartForm(20000)
		filters := formToFilter(r.Form)
		filters.Name = strings.ToLower(filters.Name)
		allResults := applyFilters(filters)

		tmpl, _ := template.New("").ParseFiles("templates/components/card-item-container.html", "templates/components/card-item.html")

		err := tmpl.ExecuteTemplate(w, "card-container", allResults)

		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		return
	}
	result, err := api.UseFallBack()
	if err != nil {
		fmt.Println(err)
	}
	data.ResultArr = result
	tmpl := template.Must(template.ParseFiles("templates/base.html", "templates/views/search.html", "templates/components/card-item.html"))
	tmpl.Execute(w, data)
	fmt.Println(len(lastRequest.AllResults))
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	lastPage = "404"
	data := Data{PageName: lastPage}

	tmpl := template.Must(template.ParseFiles("templates/base.html", "templates/views/index.html"))
	tmpl.Execute(w, data)
}

func formToFilter(form url.Values) (filters Filters) {
	filters.Name = form["name"][0]
	if _, ok := form["category"]; ok {
		filters.Category = form["category"]
	}
	if _, ok := form["mastermode"]; ok {
		filters.MasterMode = true
	}
	if _, ok := form["region"]; ok {
		filters.Regions = form["region"]
	}
	return filters
}
