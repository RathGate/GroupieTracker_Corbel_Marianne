package api

import (
	"encoding/json"
	"fmt"
	"os"
)

func GenerateFallback(normalmode bool, mastermode bool) error {
	if !mastermode && !normalmode {
		fmt.Println("WARNING: No fallback has been created since the function has only been passed false parameters.")
	}

	fullRequest, err := RequestAllEntries(true)

	// Error check:
	if err != nil {
		return err
	}

	if normalmode {
		file, err := os.OpenFile("data/all_normalmode.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
		// Error check:
		if err != nil {
			return err
		}

		json, err := json.MarshalIndent(fullRequest.NormalMode, "", "  ")
		// Error check:
		if err != nil {
			return err
		}

		_, err = file.Write(json)
		// Error check:
		if err != nil {
			return err
		}
		fmt.Println("FALLBACK: Successfully generated `/all` fallback file.")
	}

	if mastermode {
		file, err := os.OpenFile("data/all_mastermode.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
		// Error check:
		if err != nil {
			return err
		}

		json, err := json.MarshalIndent(fullRequest.MasterMode, "", "  ")
		// Error check:
		if err != nil {
			return err
		}

		_, err = file.Write(json)
		// Error check:
		if err != nil {
			return err
		}
		fmt.Println("FALLBACK: Successfully generated `/master_mode/all` fallback file.")
	}
	return nil
}

// Uses the fallback files instead of doing another request.
func UseFallBack(mastermode bool) (result []Item, err error) {
	if mastermode {
		fmt.Println("Request /master_mode/all : using fallback files to retrieve initial search data.")
	} else {
		fmt.Println("Request /all : using fallback files to retrieve initial search data.")
	}
	var filename string

	if mastermode {
		filename = "data/all_mastermode.json"
	} else {
		filename = "data/all_normalmode.json"
	}

	byteValue, err := os.ReadFile(filename)
	// Error check:
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(byteValue, &result)
	// Error check:
	if err != nil {
		return nil, err
	}

	return result, nil
}

// !DEBUG FEATURE:
func PrintAllItems(items []Item) {
	for i, v := range items {
		fmt.Printf("[%v] (%v | %v) %v\n", i, v.ID, v.Food, v.Name)
	}
}
