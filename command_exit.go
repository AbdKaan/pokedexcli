package main

import (
	"fmt"
	"os"
)

func commandExit(cfg *config, arg string) error {
	// Exit the application
	fmt.Println("Exiting Pokedex...")
	os.Exit(0)
	return nil
}
