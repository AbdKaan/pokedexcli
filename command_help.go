package main

import (
	"fmt"
	"sort"
)

func commandHelp(cfg *config, arg string) error {
	// Prints all the descriptions of the client commands available.
	cliCommands := getCommands()
	fmt.Println("\nWelcome to the Pokedex!\nUsage:")
	fmt.Println()

	// Sort the map so we get the same order every time
	keys := make([]string, 0)
	for k := range cliCommands {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// Print commands and descriptions
	for _, command := range keys {
		fmt.Printf("%s: %s\n", cliCommands[command].name, cliCommands[command].description)
	}
	fmt.Println()
	return nil
}
