package main

import (
	"errors"
	"fmt"
)

func inspect(cfg *config, name string) error {
	pokemon, ok := cfg.pokedex[name]
	if !ok {
		err := errors.New("you have not caught that pokemon")
		return err
	}

	fmt.Printf("Name: %v\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("-%v: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, typ := range pokemon.Types {
		fmt.Printf("-%v\n", typ.Type.Name)
	}

	return nil
}
