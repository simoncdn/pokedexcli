package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/simoncdn/pokedexcli/internal/pokeapi"
	"github.com/simoncdn/pokedexcli/internal/pokecache"
)

type config struct {
	pokeapiClient    pokeapi.Client
	pokeCache        *pokecache.Cache
	nextLocationsURL *string
	prevLocationsURL *string
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		userInput := scanner.Text()
		cliCommand, exist := getCommands()[userInput]

		if exist {
			err := cliCommand.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	words := strings.Fields(lowerText)

	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {
	cliCommands := map[string]cliCommand{
		"map": {
			name:        "map",
			description: "Display location areas forward",
			callback:    commandMapF,
		},
		"mapb": {
			name:        "mapb",
			description: "Display location areas back",
			callback:    commandMapB,
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
	return cliCommands
}
