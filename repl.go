package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/AbdKaan/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"map": {
			name:        "map",
			description: "Displays 20 maps on the next page",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays 20 maps on the prev page",
			callback:    commandMapb,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func startRepl(cfg *config) {
	// Create scanner and the infinite loop
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("pokedex > ")
		// Read input
		scanner.Scan()
		input := scanner.Text()
		clean_input := cleanInput(input)
		if len(clean_input) == 0 {
			continue
		}

		// Clean the input and call the command
		commandName := clean_input[0]
		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg)
			if err != nil {
				err = fmt.Errorf("error trying to execute the command: %w", err)
				fmt.Println(err)
			}
		} else {
			fmt.Println("Invalid command.")
		}
	}
}
