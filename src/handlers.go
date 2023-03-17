package main

import (
	"fmt"
	"groupie-tracker/packages/api"
	"log"
	"net/http"
	"net/url"
	"regexp"
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

	if _, err := strconv.Atoi(id); err != nil {
		notFoundHandler(w, r)
		return
	}

	reg := regexp.MustCompile(`\?mode=master`)
	var request api.Item
	var err error
	if reg.MatchString(r.URL.RequestURI()) {
		request, err = api.MakeEntryRequest(id, true)
	} else {
		request, err = api.MakeEntryRequest(id, false)
	}

	if err != nil || request.Name == "" {
		notFoundHandler(w, r)
		return
	}
	data := Data{PageName: lastPage, PerfectMatch: request}
	tmpl, err := template.New("base.html").Funcs(template.FuncMap{
		"addPaddingToNumber": addPaddingToNumber,
	}).ParseFiles("templates/base.html", "templates/views/item.html", "templates/components/entry-item.html")
	if err != nil {
		log.Fatal(err)
	}
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

			tmpl, err := template.New("").Funcs(template.FuncMap{
				"addPaddingToNumber": addPaddingToNumber,
			}).ParseFiles("templates/components/card-item-container.html", "templates/components/card-item.html")
			if err != nil {
				log.Fatal(err)
			}
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
	tmpl, err := template.New("base.html").Funcs(template.FuncMap{
		"addPaddingToNumber": addPaddingToNumber,
	}).ParseFiles("templates/base.html", "templates/views/categories.html", "templates/components/card-item.html")
	if err != nil {
		log.Fatal(err)
	}
	tmpl.Execute(w, data)

	lastRequest.updateRequest("idk", temp)
}
func searchHandler(w http.ResponseWriter, r *http.Request) {
	data := Data{PageName: "search", Regions: REGIONS_NAMES, Categories: damn}
	if r.Method == "POST" {
		r.ParseMultipartForm(20000)
		filters := formToFilter(r.Form)

		allResults := applyFilters(filters)

		tmpl, err := template.New("").Funcs(template.FuncMap{
			"addPaddingToNumber": addPaddingToNumber,
		}).ParseFiles("templates/components/card-item-container.html", "templates/components/card-item.html")
		if err != nil {
			log.Fatal(err)
		}
		err = tmpl.ExecuteTemplate(w, "card-container", allResults)

		if err != nil {
			log.Fatal(err)
		}
		return
	}
	result, err := api.UseFallBack(false)
	if err != nil {
		fmt.Println(err)
	}
	data.ResultArr = result

	// files := []string{"templates/_results-table.html", "templates/results.html"}
	tmpl, err := template.New("base.html").Funcs(template.FuncMap{
		"addPaddingToNumber": addPaddingToNumber,
	}).ParseFiles("templates/base.html", "templates/views/search.html", "templates/components/card-item.html")
	if err != nil {
		panic(err)
	}

	tmpl.Execute(w, data)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)

	tmpl := template.Must(template.ParseFiles("templates/base.html", "templates/views/404.html"))
	tmpl.Execute(w, nil)
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
