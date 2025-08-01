package main

import "fmt"

func commandHelp(cfg *config, args ...string) error {
	cliCommands := getCommands()

	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, cliCommand := range cliCommands {
		fmt.Printf("%s: %s\n", cliCommand.name, cliCommand.description)
	}
	fmt.Println()
	return nil
}
