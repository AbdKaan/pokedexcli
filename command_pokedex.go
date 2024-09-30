package main

import "fmt"

func pokedex(cfg *config, arg string) error {
	if len(cfg.pokedex) == 0 {
		fmt.Println("You haven't caught any pokemons yet")
		return nil
	}

	for _, pokemon := range cfg.pokedex {
		fmt.Printf(" - %v\n", pokemon.Name)
	}

	return nil
}
