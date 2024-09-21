package main

import (
	"fmt"
	"os"
)

func commandExit() error {
	// Exit the application
	fmt.Println("Exiting Pokedex...")
	os.Exit(0)
	return nil
}
