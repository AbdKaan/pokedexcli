package main

import (
	"encoding/json"
	"fmt"

	"github.com/AbdKaan/pokedexcli/internal/pokeapi"
)

func explore(cfg *config, area string) error {
	if area == "" {
		return fmt.Errorf("haven't provided area to explore")
	}

	var details pokeapi.Area
	var err error

	data, found := cfg.pokeCache.Get(area)
	if !found {
		details, err = cfg.pokeapiClient.DetailLocation(cfg.pokeCache, area)
		if err != nil {
			return fmt.Errorf("%w, are you sure the area you entered exists?", err)
		}
	} else {
		if err := json.Unmarshal(data, &details); err != nil {
			return err
		}
	}

	fmt.Printf("Exploring %s...\nFound Pokemon:\n", area)
	for _, encounter := range details.PokemonEncounters {
		fmt.Printf("- %s\n", encounter.Pokemon.Name)
	}
	return nil
}
