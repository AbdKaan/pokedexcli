package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/AbdKaan/pokedexcli/internal/pokeapi"
)

func commandMapf(cfg *config) error {
	// Check if the location is in cache, if it's not in cache then call ListLocations
	// which will find the next locations and add them to cache
	locationsResp := pokeapi.LocationArea{}

	// Check if nextLocationsURL is initialized to a pointer
	var result []byte
	var found bool
	if cfg.nextLocationsURL == nil {
		result, found = cfg.pokeCache.Get("")
	} else {
		result, found = cfg.pokeCache.Get(*cfg.nextLocationsURL)
	}

	if !found {
		var err error
		locationsResp, err = cfg.pokeapiClient.ListLocations(cfg.pokeCache, cfg.nextLocationsURL)
		if err != nil {
			return fmt.Errorf("error trying to get locations: %w", err)
		}
	} else {
		if err := json.Unmarshal(result, &locationsResp); err != nil {
			return err
		}
	}

	cfg.nextLocationsURL = &locationsResp.Next
	cfg.prevLocationsURL = &locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Printf("- %s\n", loc.Name)
	}
	return nil
}

func commandMapb(cfg *config) error {
	if cfg.prevLocationsURL == nil || *cfg.prevLocationsURL == "" {
		return errors.New("you're on the first page")
	}

	// Check if the location is in cache, if it's not in cache then call ListLocations
	// which will find the next locations and add them to cache
	locationsResp := pokeapi.LocationArea{}
	if result, found := cfg.pokeCache.Get(*cfg.prevLocationsURL); !found {
		var err error
		locationsResp, err = cfg.pokeapiClient.ListLocations(cfg.pokeCache, cfg.prevLocationsURL)
		if err != nil {
			return fmt.Errorf("error trying to get locations: %w", err)
		}
	} else {
		if err := json.Unmarshal(result, &locationsResp); err != nil {
			return err
		}
	}

	cfg.nextLocationsURL = &locationsResp.Next
	cfg.prevLocationsURL = &locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Printf("- %s\n", loc.Name)
	}
	return nil
}
