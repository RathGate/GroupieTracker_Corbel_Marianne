package main

import (
	"fmt"
	"groupie-tracker/packages/api"
	"net/http"
	"regexp"
	"strconv"
	"text/template"

	"github.com/gorilla/mux"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	data := Data{PageName: "Home"}

	tmpl := generateTemplate("base.html", []string{"templates/base.html", "templates/views/index.html"})
	tmpl.Execute(w, data)
}

func itemHandler(w http.ResponseWriter, r *http.Request) {
	data := Data{PageName: "Items"}
	var request api.Item
	var err error

	// Retrieves Item ID from URL
	id := mux.Vars(r)["id"]

	// Error check || Sends to 404 page if user entered an invalid ID
	if _, err := strconv.Atoi(id); err != nil {
		notFoundHandler(w, r)
		return
	}

	// Checks if ID comes from normal mode or mastermode based on URL parameter
	if regexp.MustCompile(`\?mode=master`).MatchString(r.URL.RequestURI()) {
		request, err = api.RequestSingleEntry(id, true)
	} else {
		request, err = api.RequestSingleEntry(id, false)
	}
	// Error check || Sends to 404 if the item's name is empty (== no item):
	if err != nil || request.Name == "" {
		notFoundHandler(w, r)
		return
	}

	data.PerfectMatch = request

	// *Generates and executes templates:
	tmpl := generateTemplate("base.html", []string{"templates/base.html", "templates/views/item.html", "templates/components/entry-item.html"})
	tmpl.Execute(w, data)
}

func categoriesHandler(w http.ResponseWriter, r *http.Request) {
	data := Data{PageName: "Categories"}

	// ?User came here by POST method:
	if r.Method == "POST" {
		var categoryEntries []api.Item

		// Retrieves category name from form:
		if categoryStr := r.FormValue("category-id"); categoryStr != "" {

			categoryID, err := strconv.Atoi(categoryStr)
			if err != nil || categoryID < 0 || categoryID > 4 {
				notFoundHandler(w, r)
				return
			}

			// Makes the request to the API:
			categoryEntries, err = api.RequestEntriesByCategory(CREATURES_NAMES[categoryID])
			// Error check:
			if err != nil {
				fmt.Println("WARNING: Something went wrong with the API. It might be currently offline.")
				w.WriteHeader(404)
				return
			}

			// Updates last request value:
			lastRequest.updateRequest(data.PageName, categoryEntries)

			// *Generates and executes templates:
			tmpl := generateTemplate("", []string{"templates/components/card-item-container.html", "templates/components/card-item.html"})
			tmpl.ExecuteTemplate(w, "card-container", categoryEntries)
			return
		}
	}

	// ?User came here by GET method:
	categoryEntries, err := api.RequestEntriesByCategory("creatures")
	if err != nil {
		println(err)
		return
	}

	data.ResultArr = categoryEntries

	// Updates last request value:
	lastRequest.updateRequest(data.PageName, categoryEntries)

	// *Generates and executes templates:
	tmpl := generateTemplate("base.html", []string{"templates/base.html", "templates/views/categories.html", "templates/components/card-item.html"})
	tmpl.Execute(w, data)
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	data := Data{PageName: "Search", Regions: REGIONS_NAMES, Categories: CREATURES_NAMES}

	// ?User came here by "POST" method:
	if r.Method == "POST" {
		// Parses the form and creates the search filters:
		r.ParseMultipartForm(20000)
		filters := api.FormToFilter(r.Form)

		// Generates the results based on the filters:
		resultItems := api.ApplyFilters(filters)

		// *Generates and executes templates:
		tmpl := generateTemplate("", []string{"templates/components/card-item-container.html", "templates/components/card-item.html"})
		tmpl.ExecuteTemplate(w, "card-container", resultItems)
		return
	}

	// ?User came here by "GET" method:
	resultItems, err := api.UseFallBack(false)
	// Error check
	if err != nil {
		fmt.Println("Someting went wrong with the fallback files.")
	}

	data.ResultArr = resultItems

	// *Generates and executes templates:
	tmpl := generateTemplate("base.html", []string{"templates/base.html", "templates/views/search.html", "templates/components/card-item.html"})
	tmpl.Execute(w, data)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)

	// *Generates and executes templates:
	tmpl := template.Must(template.ParseFiles("templates/base.html", "templates/views/404.html"))
	tmpl.Execute(w, nil)
}
