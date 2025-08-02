package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}

	locationName := args[0]
	locationRes, err := cfg.pokeapiClient.GetLocation(locationName)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", locationRes.Name)
	fmt.Println("Found Pokemon: ")

	for _, enc := range locationRes.PokemonEncounters {
		fmt.Println("- " + enc.Pokemon.Name)
	}

	return nil
}
