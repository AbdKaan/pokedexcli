package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
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

func startRepl() {
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
			err := command.callback()
			if err != nil {
				fmt.Errorf("Error trying to execute the command: %w\n", err)
			}
		} else {
			fmt.Println("Invalid command.")
		}
	}
}
