package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

func commandCatch(cfg *config, args []string) error {
	if len(args) != 1 {
		return errors.New("catch takes one pokemon name arg")
	}

	name := args[0]
	pokemonInfo, err := cfg.pokeapiClient.GetPokemonInfo(name)
	if err != nil {
		return fmt.Errorf("couldn't fetch pokemonInfo for %s. Error: %w", name, err)
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	// Handle random catching
	catchProb := 1.0 - (float64(pokemonInfo.BaseExperience) / 1000)
	randFloat := rand.Float64()
	randFloat = float64(int(randFloat*100)) / 100
	if randFloat > catchProb {
		fmt.Printf("%s escaped!\n", name)
		return nil
	}
	fmt.Printf("%s was caught!\n", name)

	//Add if not in pokedex
	if _, ok := cfg.Pokedex[name]; !ok {
		cfg.Pokedex[name] = pokemonInfo
	}

	return nil

}
