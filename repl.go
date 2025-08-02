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
	pokedex          map[string]pokeapi.Pokemon
	nextLocationsURL *string
	prevLocationsURL *string
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		userInput := cleanInput(scanner.Text())

		cliCommand, exist := getCommands()[userInput[0]]
		args := []string{}
		if len(userInput) > 1 {
			args = userInput[1:]
		}

		if exist {
			err := cliCommand.callback(cfg, args...)
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
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	cliCommands := map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Show all of your Pokemon in your Pokedex",
			callback:    commandPokedex,
		},
		"inspect": {
			name:        "inspect <pokemon_name>",
			description: "View details about a caught Pokemon (stats, types, etc...)",
			callback:    commandInspect,
		},
		"catch": {
			name:        "catch <pokemon_name>",
			description: "Try to catch a Pokemon",
			callback:    commandCatch,
		},
		"explore": {
			name:        "explore <location_name>",
			description: "Explore a location",
			callback:    commandExplore,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMapF,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapB,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
	return cliCommands
}
