package main

import (
	"encoding/json"
	"groupie-tracker/packages/api"
	"os"
)

func GenerateFallback() error {
	result, err := api.MakeFullRequest()
	if err != nil {
		return err
	}
	file, err := json.MarshalIndent(result, "", "   ")
	if err != nil {
		return err
	}
	err = os.WriteFile("/assets/data/fallback.json", file, 0644)
	return err
}
