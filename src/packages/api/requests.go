package api

import (
	"encoding/json"
	"io"
	"net/http"
	"sort"
)

func MakeFullRequest(mastermode bool) (result []Item, err error) {
	var temp FullRequest
	var accessory CategoryRequest
	if mastermode {
		reqBody, status, err := MakeRequest("https://botw-compendium.herokuapp.com/api/v2/master_mode/all")

		if status != http.StatusOK {
			return []Item{}, nil
		}
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(reqBody, &accessory)
		if err != nil {
			return nil, err
		}
	}

	reqBody, status, err := MakeRequest("https://botw-compendium.herokuapp.com/api/v2/all")

	if status != http.StatusOK {
		return []Item{}, nil
	}
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(reqBody, &temp)
	if err != nil {
		return nil, err
	}
	temp.Data.Monsters = append(temp.Data.Monsters, accessory.Items...)
	return FlattenFullRequest(temp), err
}

func MakeCategoryRequest(categoryName string) ([]Item, error) {
	reqBody, status, err := MakeRequest("https://botw-compendium.herokuapp.com/api/v2/category/" + categoryName)

	if status != http.StatusOK {
		return []Item{}, nil
	}
	if err != nil {
		return nil, err
	}

	if categoryName == "creatures" {
		var result CreaturesRequest
		err = json.Unmarshal(reqBody, &result)
		return FlattenCreatureRequest(result), err
	}

	var result CategoryRequest

	err = json.Unmarshal(reqBody, &result)
	sort.Slice(result.Items, func(a, b int) bool {
		return result.Items[a].ID < result.Items[b].ID
	})
	return result.Items, err
}

func MakeEntryRequest(entryName string, mastermode bool) (result Item, err error) {
	var requestURL string
	if mastermode {
		requestURL = "https://botw-compendium.herokuapp.com/api/v2/master_mode/entry/" + entryName
	} else {
		requestURL = "https://botw-compendium.herokuapp.com/api/v2/entry/" + entryName
	}
	reqBody, status, err := MakeRequest(requestURL)
	var temp EntryRequest
	if status != http.StatusOK {
		return result, nil
	}
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(reqBody, &temp)
	return temp.Item, err
}

func MakeRequest(url string) (result []byte, status int, err error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)

	res, err := client.Do(req)
	if err != nil {
		return nil, http.StatusOK, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, res.StatusCode, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, http.StatusOK, err
	}
	return body, http.StatusOK, nil
}
