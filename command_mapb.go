package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func commandMapb() error {
	// Check if we are on first 2 pages
	if locationArea.Next == "https://pokeapi.co/api/v2/location-area?offset=20&limit=20" || locationArea.Next == "https://pokeapi.co/api/v2/location-area" {
		noPrevErr := errors.New("no previous pages exist")
		return noPrevErr
	}

	// Step 1: Make the GET request
	resp, err := http.Get(locationArea.Previous)
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
