package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandMap() error {
	// Step 1: Make the GET request
	resp, err := http.Get(locationArea.Next)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Step 2: Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Step 3: Unmarshal the JSON data into the struct
	if err := json.Unmarshal(body, &locationArea); err != nil {
		return err
	}

	// Step 4: Print the result
	for _, result := range locationArea.Results {
		fmt.Println(result.Name)
	}
	return nil
}
