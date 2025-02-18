package main

import (
	"time"

	"github.com/AbdKaan/pokedexcli/internal/pokeapi"
	"github.com/AbdKaan/pokedexcli/internal/pokecache"
)

func main() {
	// Create client
	pokeClient := pokeapi.NewClient(5 * time.Second)

	// Create cache
	pokeCache := pokecache.NewCache(5 * time.Minute)

	pokedex := make(map[string]pokeapi.Pokemon)

	// Create config
	cfg := &config{
		pokeapiClient: pokeClient,
		pokeCache:     pokeCache,
		pokedex:       pokedex,
	}

	startRepl(cfg)
}
