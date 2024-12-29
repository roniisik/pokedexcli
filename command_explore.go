package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args []string) error {
	if len(args) != 1 {
		return errors.New("explore takes one area name arg")
	}

	areaInfo, err := cfg.pokeapiClient.ListPokemons(args[0])
	if err != nil {
		fmt.Printf("Unable to retrieve areaInfo, error:\n")
		return err
	}

	for _, enc := range areaInfo.PokemonEncounters {
		fmt.Println(enc.Pokemon.Name)
	}

	return nil
}
