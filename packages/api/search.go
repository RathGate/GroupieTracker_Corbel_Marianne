package api

import (
	"encoding/json"
	"os"
	"reflect"
	"sort"
	"strings"
)

func SearchByName(useFallback bool, searchName string, excludedId int) (result []Item, err error) {
	var temp []Item
	if !useFallback {
		temp, err = MakeFullRequest()
		if err != nil {
			return nil, err
		}
	} else {
		temp, err = UseFallBack()
		if err != nil {
			return nil, err
		}
	}
	for _, element := range temp {
		if strings.Contains(element.Name, searchName) && element.ID != excludedId {
			result = append(result, element)
		}
	}
	return result, nil
}

func UseFallBack() (temp []Item, err error) {
	byteValue, err := os.ReadFile("assets/data/fallback.json")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(byteValue, &temp)
	if err != nil {
		return nil, err
	}
	return temp, nil
}

// TODO: Clear the function (redundancy)
// Flattens the multilevel FullRequest struct (see structs.go)
// to a one-dimension []Item array.
func FlattenFullRequest(request FullRequest) (resultArr []Item) {
	var temp = request.Data

	// Flatten 2D struct Creatures.Food and Creatures.NonFood
	e := reflect.ValueOf(&temp.Creatures).Elem()
	for i := 0; i < e.NumField(); i++ {
		arr := (e.Field(i).Interface().([]Item))
		for j := 0; i == 0 && j < len(arr); j++ {
			arr[j].Food = true
		}
		resultArr = append(resultArr, arr...)
	}

	e = reflect.ValueOf(&temp).Elem()
	for i := 1; i < e.NumField(); i++ {
		resultArr = append(resultArr, e.Field(i).Interface().([]Item)...)
	}
	sort.Slice(resultArr, func(a, b int) bool {
		return resultArr[a].ID < resultArr[b].ID
	})
	return resultArr
}
