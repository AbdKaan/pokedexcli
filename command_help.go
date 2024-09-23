package main

import "fmt"

func commandHelp(cfg *config) error {
	// Prints all the descriptions of the client commands available.
	cliCommands := getCommands()
	fmt.Println("\nWelcome to the Pokedex!\nUsage:")
	fmt.Println()
	for _, command := range cliCommands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	fmt.Println()
	return nil
}
