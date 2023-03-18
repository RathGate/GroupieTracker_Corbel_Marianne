package utils

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

// ?Sends a GET request to the specified URL. No optional parameters can be passed.
func MakeRequest(url string) (result []byte, status int, err error) {
	client := &http.Client{}

	req, _ := http.NewRequest("GET", url, nil)
	res, err := client.Do(req)

	// Error check :
	if err != nil || res.StatusCode != http.StatusOK {
		return nil, res.StatusCode, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	return body, http.StatusOK, err
}

// ?Generates an Error if http.Status is not correct.
func GenerateError(err error, status int) error {
	if status != http.StatusOK {
		return fmt.Errorf("server responded with a %v status", status)
	}
	return err
}

// ?Adds padding "0"s to create a 3-digit number.
func AddPaddingToNumber(baseInt int) string {
	baseStr := strconv.Itoa(baseInt)
	if len(baseStr) >= 3 {
		return baseStr
	}
	return fmt.Sprintf("%v%v", strings.Repeat("0", 3-len(baseStr)), baseStr)
}

// ?Checks if string is present in slice of strings.
func StringInSlice(a string, list []string) bool {
	if len(list) == 0 {
		return true
	}
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
