package api

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/packages/utils"
	"net/url"
	"os"
	"strings"
)

type Filters struct {
	Name       string
	Regions    []string
	Category   []string
	MasterMode bool
}

type Places struct {
	All           []string `json:"All"`
	Akkala        []string `json:"Akkala"`
	CentralHyrule []string `json:"Central Hyrule"`
	Eldin         []string `json:"Eldin"`
	Faron         []string `json:"Faron"`
	Gerudo        []string `json:"Gerudo"`
	Hebra         []string `json:"Hebra"`
	Lanayru       []string `json:"Lanayru"`
	Necluda       []string `json:"Necluda"`
}

func LoadPlaceNames() (temp map[string][]string, err error) {
	byteValue, err := os.ReadFile("data/places.json")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(byteValue, &temp)
	if err != nil {
		return nil, err
	}
	return temp, nil
}

func FormToFilter(form url.Values) (filters Filters) {

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

func IsInRegion(areas []string, allregions map[string][]string, item Item) bool {
	if len(areas) == 0 || len(item.CommonLocations) == 0 || utils.StringInSlice("Greater Hyrule", item.CommonLocations) {
		return true
	}
	for _, area := range areas {

		region, ok := allregions[area]
		if !ok {
			continue
		} else {
			for _, location := range item.CommonLocations {
				if utils.StringInSlice(location, region) {
					return true
				}
			}
		}
	}
	return false
}

func ApplyFilters(filters Filters) (result []Item) {
	allregions, _ := LoadPlaceNames()

	var allitems []Item
	var err error

	if filters.MasterMode {
		allitems, _ = UseFallBack(true)
	} else {
		allitems, err = UseFallBack(false)
	}
	// Error check:
	if err != nil {
		fmt.Println(err)
	}

	filters.Name = strings.ToLower(filters.Name)

	for _, item := range allitems {
		if strings.Contains(item.Name, filters.Name) && IsInRegion(filters.Regions, allregions, item) && utils.StringInSlice(item.Category, filters.Category) {
			result = append(result, item)
		}
	}

	return result
}
