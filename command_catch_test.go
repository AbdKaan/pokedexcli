package main

import (
	"strings"
	"testing"
	"time"

	"github.com/AbdKaan/pokedexcli/internal/pokeapi"
)

func TestCatchPokemon(t *testing.T) {
	// Create client
	pokeClient := pokeapi.NewClient(5 * time.Second)

	// Create pokedex
	pokedex := make(map[string]pokeapi.Pokemon)

	// Create config
	cfg := &config{
		pokeapiClient: pokeClient,
		pokedex:       pokedex,
	}

	cases := []string{
		"pikachu",
		"bulbasaur",
	}

	for _, c := range cases {
		catch(cfg, c)
		if strings.ToLower(cfg.pokedex[c].Name) != c {
			t.Errorf("expected to find pikachu, found %s", cfg.pokedex[c].Name)
			return
		}
	}
}
