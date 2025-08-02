package main

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/simoncdn/pokedexcli/internal/pokeapi"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	pokemonName := args[0]

	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	source := rand.NewSource(int64(pokemon.BaseExperience))
	rng := rand.New(source)

	catchChance := rng.Float64()
	difficulty := float64(pokemon.BaseExperience) / 300.0

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	fmt.Println()
	if catchChance > difficulty {
		fmt.Printf("%s was caught!\n", pokemonName)
		addToPokedex(cfg, pokemon)
		fmt.Println("You may now inspect it with the inspect command.")
	} else {
		fmt.Printf("%s espaced!\n", pokemonName)
	}

	return nil
}

func addToPokedex(cfg *config, pokemon pokeapi.Pokemon) {
	cfg.pokedex[pokemon.Name] = pokemon
	fmt.Printf("%v added to the Pokedex!\n", pokemon.Name)
}
