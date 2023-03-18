package api

import (
	"encoding/json"
	"groupie-tracker/packages/utils"
	"net/http"
	"reflect"
	"sort"
)

type FullModeRequest struct {
	NormalMode []Item
	MasterMode []Item
}

// Sends requests to /all and/or /master_mode/all,
// processes and sorts the data as well.
func RequestAllEntries(mastermode bool) (result FullModeRequest, err error) {
	var normalResponse FullRequest

	reqBody, status, err := utils.MakeRequest("https://botw-compendium.herokuapp.com/api/v2/all")
	// Error check:
	if status != http.StatusOK || err != nil {
		return FullModeRequest{}, utils.GenerateError(err, status)
	}

	err = json.Unmarshal(reqBody, &normalResponse)
	// Error check
	if err != nil {
		return FullModeRequest{}, err
	}

	result.NormalMode = FlattenFullRequest(normalResponse)

	if !mastermode {
		return result, nil
	}

	result.MasterMode, err = RequestMastermodeEntries(result.NormalMode)
	return result, err
}

func RequestMastermodeEntries(baseEntries []Item) (result []Item, err error) {
	if len(baseEntries) == 0 {
		req, err := RequestAllEntries(false)
		// Error check:
		if err != nil {
			return nil, err
		}
		baseEntries = req.NormalMode
	}

	var masterResponse CategoryRequest

	reqBody, status, err := utils.MakeRequest("https://botw-compendium.herokuapp.com/api/v2/master_mode/all")
	// Error check:
	if status != http.StatusOK || err != nil {
		return nil, utils.GenerateError(err, status)
	}

	err = json.Unmarshal(reqBody, &masterResponse)
	// Error check
	if err != nil {
		return nil, err
	}

	masterItems := masterResponse.Items
	return insertItemsByID(baseEntries, masterItems), nil
}

// Sends requests to /category/{categoryName},
// processes and sorts the data as well.
func RequestEntriesByCategory(categoryName string) ([]Item, error) {
	reqBody, status, err := utils.MakeRequest("https://botw-compendium.herokuapp.com/api/v2/category/" + categoryName)
	// Error check:
	if status != http.StatusOK || err != nil {
		return nil, utils.GenerateError(err, status)
	}

	// `creatures` category is a bit specific since the items will be written
	// one layer deeper in the JSON response: it'll need to be flattened.
	if categoryName == "creatures" {
		var result CreaturesRequest

		err = json.Unmarshal(reqBody, &result)
		// Error check:
		if err != nil {
			return nil, err
		}

		return FlattenCreatureRequest(result), err
	}

	var result CategoryRequest

	err = json.Unmarshal(reqBody, &result)
	if err != nil {
		return nil, err
	}

	return SortItemsByID(result.Items), err
}

// Sends requests to /entry/{entryName} or /master_mode/entry/{entryName}
//
//	depending on if `mastermode` is true of not.
func RequestSingleEntry(entryName string, mastermode bool) (Item, error) {

	// Formats the URL according to the chosen mode:
	var requestURL string
	if mastermode {
		requestURL = "https://botw-compendium.herokuapp.com/api/v2/master_mode/entry/" + entryName
	} else {
		requestURL = "https://botw-compendium.herokuapp.com/api/v2/entry/" + entryName
	}

	reqBody, status, err := utils.MakeRequest(requestURL)
	// Error check:
	if status != http.StatusOK || err != nil {
		return Item{}, utils.GenerateError(err, status)
	}

	var result EntryRequest

	err = json.Unmarshal(reqBody, &result)
	// Error check:
	if err != nil {
		return Item{}, err
	}

	return result.Item, err
}

// Fuses an array of Items into another.
// Caution: each Item of newItems will be insert in baseItems by their IDs.
func insertItemsByID(originalArr []Item, newItems []Item) []Item {
	newItems = SortItemsByID(newItems)
	var baseItems []Item
	baseItems = append(baseItems, SortItemsByID(originalArr)...)

	for _, item := range newItems {
		item.MasterExclusive = true
		baseItems = append(baseItems[:item.ID], baseItems[item.ID-1:]...)
		baseItems[item.ID-1] = item
	}

	for i := 0; i < len(baseItems); i++ {
		baseItems[i].MasterID = i + 1
		baseItems[i].DisplayMaster = true
	}
	return baseItems
}

// Sorts an array of Items by their IDs.
func SortItemsByID(items []Item) []Item {
	sort.Slice(items, func(a, b int) bool {
		return items[a].ID < items[b].ID
	})
	return items
}

// TODO (notes for later): Both FlattenFullRequest and FlattenCreatureRequest present some redundancy.
// TODO: Maybe possible to factorize the functions to a recursive one using maps or interfaces?
// Flattens multidimensional FullRequest struct.
func FlattenFullRequest(request FullRequest) (resultArr []Item) {
	var temp = request.Data

	// Flattens 2D struct Creatures.Food and Creatures.NonFood
	e := reflect.ValueOf(&temp.Creatures).Elem()

	for i := 0; i < e.NumField(); i++ {
		arr := (e.Field(i).Interface().([]Item))
		// Adds "Food" = "true" to all items from field "Food"
		if e.Type().Field(i).Name == "Food" {
			arr := (e.Field(i).Interface().([]Item))
			for j := 0; j < len(arr); j++ {
				arr[j].Food = true
			}
		}
		// Appends the nested Items
		resultArr = append(resultArr, arr...)
	}

	// Flattens all the other 1D structs:
	e = reflect.ValueOf(&temp).Elem()
	for i := 1; i < e.NumField(); i++ {
		resultArr = append(resultArr, e.Field(i).Interface().([]Item)...)
	}

	return SortItemsByID(resultArr)
}

// Flattens 2-dimension CreatureRequest struct (Items received from API are originally)
// separated in two
func FlattenCreatureRequest(request CreaturesRequest) (resultArr []Item) {

	var temp = request.Data

	e := reflect.ValueOf(&temp).Elem()

	for i := 0; i < e.NumField(); i++ {
		// Adds "Food" = "true" to all items from field "Food"
		if e.Type().Field(i).Name == "Food" {
			arr := (e.Field(i).Interface().([]Item))
			for j := 0; j < len(arr); j++ {
				arr[j].Food = true
			}
		}
		// Appends the nested Items
		resultArr = append(resultArr, e.Field(i).Interface().([]Item)...)
	}

	return SortItemsByID(resultArr)
}
