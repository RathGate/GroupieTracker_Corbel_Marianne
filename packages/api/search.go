package api

import (
	"reflect"
	"sort"
)

// func SearchByName(searchName string) {
// 	full, err := MakeFullRequest()
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// }

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
