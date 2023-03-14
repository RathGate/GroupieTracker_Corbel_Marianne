package main

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/packages/api"
	"net/http"
	"os"
	"strings"
)

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

func GenerateFallback() error {
	result, err := api.MakeFullRequest(false)
	if err != nil {
		return err
	}

	file, err := json.MarshalIndent(result, "", "   ")
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = os.WriteFile("assets/data/places.json", file, 0644)
	fmt.Println(err)
	return err
}

func GenerateMMFallback() {
	baseArr, _ := api.MakeFullRequest(false)

	var masterData api.CategoryRequest
	reqBody, status, _ := api.MakeRequest("https://botw-compendium.herokuapp.com/api/v2/master_mode/all")
	if status != http.StatusOK {
		return
	}
	_ = json.Unmarshal(reqBody, &masterData)
	for _, item := range masterData.Items {
		baseArr = append(baseArr[:item.ID], baseArr[item.ID-1:]...)
		baseArr[item.ID-1] = item
	}

	for i := 0; i < len(baseArr); i++ {
		baseArr[i].ID = i + 1
		fmt.Printf("%v | %v\n", baseArr[i].ID, baseArr[i].Name)
	}
	file, err := json.MarshalIndent(baseArr, "", "   ")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = os.WriteFile("assets/data/mastermode.json", file, 0644)
	fmt.Println(err)
}

func Fill() (temp map[string][]string, err error) {
	byteValue, err := os.ReadFile("assets/data/places.json")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(byteValue, &temp)
	if err != nil {
		return nil, err
	}
	return temp, nil
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func isInRegion(areas []string, item api.Item) bool {
	allregions, _ := Fill()
	if len(areas) == 0 {
		return true
	}
	for _, area := range areas {
		for _, location := range item.CommonLocations {
			if stringInSlice(location, allregions[area]) {
				return true
			}
		}
	}
	return false
}

type Filters struct {
	Name       string
	Regions    []string
	Category   []string
	MasterMode bool
}

func applyFilters(filters Filters) (result []api.Item) {
	var allitems []api.Item
	var err error
	if !filters.MasterMode {
		allitems, _ = api.UseFallBack()
	} else {
		allitems, err = api.MakeFullRequest(true)
		if err != nil {
			fmt.Println(err)
		}
	}

	for _, item := range allitems {
		if strings.Contains(item.Name, filters.Name) && isInRegion(filters.Regions, item) && stringInSlice(item.Category, filters.Category) {
			result = append(result, item)
		}
	}
	return result
}
