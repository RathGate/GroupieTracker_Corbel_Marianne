package main

import (
	"groupie-tracker/packages/api"
	"groupie-tracker/packages/utils"
	"log"
	"text/template"
)

// ?GLOBAL VARIABLES:

var REGIONS_NAMES = []string{"Akkala", "Central Hyrule", "Eldin", "Faron", "Gerudo", "Hebra", "Lanayru", "Necluda"}
var CREATURES_NAMES = []string{"creatures", "monsters", "materials", "equipment", "treasure"}

var lastRequest LastRequest

// ?STORING AND UPDATING LAST REQUEST:

type LastRequest struct {
	Page       string
	AllResults []api.Item
}

func (request *LastRequest) updateRequest(page string, results []api.Item) {
	lastRequest.Page = page
	lastRequest.AllResults = results
}

// ?TEMPLATES FUNCTIONS:

func generateTemplate(templateName string, filepaths []string) *template.Template {
	tmpl, err := template.New(templateName).Funcs(template.FuncMap{
		"addPaddingToNumber": utils.AddPaddingToNumber,
	}).ParseFiles(filepaths...)
	// Error check:
	if err != nil {
		log.Fatal(err)
	}
	return tmpl
}

// ?TEMPLATES STRUCT:

// Data sent to the templates:
type Data struct {
	PageName     string
	DataType     string
	PerfectMatch api.Item
	ResultArr    []api.Item
	Regions      []string
	Categories   []string
	Mastermode   bool
}
