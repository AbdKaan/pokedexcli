package main

import (
	"fmt"
	"math/rand"
)

func catch(cfg *config, pokemon string) error {
	if pokemon == "" {
		return fmt.Errorf("haven't provided a pokemon to catch")
	}

	pokemonDetails, err := cfg.pokeapiClient.DetailPokemon(pokemon)
	if err != nil {
		return fmt.Errorf("%w, are you sure the pokemon you entered exists?", err)
	}

	rng := rand.Intn(100)
	pokemonBaseExp := pokemonDetails.BaseExperience

	if pokemonBaseExp < 100 {
		if rng < 75 {
			fmt.Printf("Congratulations, you caught %s!\n", pokemonDetails.Name)
			cfg.pokedex[pokemon] = pokemonDetails
		} else {
			fmt.Println("Failed to catch. Git Gud!")
		}
	} else if pokemonBaseExp < 200 {
		if rng < 40 {
			fmt.Printf("Congratulations, you caught %s!\n", pokemonDetails.Name)
			cfg.pokedex[pokemon] = pokemonDetails
		} else {
			fmt.Println("Failed to catch. Git Gud!")
		}
	} else {
		if rng < 15 {
			fmt.Printf("Congratulations, you caught %s!\n", pokemonDetails.Name)
			cfg.pokedex[pokemon] = pokemonDetails
		} else {
			fmt.Println("Failed to catch. Git Gud!")
		}
	}

	return nil
}
