package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/AbdKaan/pokedexcli/internal/pokeapi"
	"github.com/AbdKaan/pokedexcli/internal/pokecache"
)

type config struct {
	pokeapiClient    pokeapi.Client
	pokeCache        *pokecache.Cache
	pokedex          map[string]pokeapi.Pokemon
	nextLocationsURL *string
	prevLocationsURL *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"pokedex": {
			name:        "pokedex",
			description: "Displays the pokemons in your pokedex",
			callback:    pokedex,
		},
		"catch": {
			name:        "catch",
			description: "Try to catch a pokemon! (catch <pokemon>)",
			callback:    catch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a pokemon in your pokedex (inspect <pokemon>)",
			callback:    inspect,
		},
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
		"explore": {
			name:        "explore",
			description: "Displays pokemons in a given area, need to provide the area as an argument (explore <area>)",
			callback:    explore,
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
		fmt.Print("Pokedex > ")
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
			var err error
			if len(clean_input) > 1 {
				err = command.callback(cfg, clean_input[1])
			} else {
				err = command.callback(cfg, "")
			}

			if err != nil {
				err = fmt.Errorf("error trying to execute the command: %w", err)
				fmt.Println(err)
			}
		} else {
			fmt.Println("Invalid command.")
		}
	}
}
