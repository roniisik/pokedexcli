package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args []string) error {
	if len(args) != 1 {
		return errors.New("catch takes one pokemon name arg")
	}

	name := args[0]
	pokemonInfo, err := cfg.pokeapiClient.GetPokemonInfo(name)
	if err != nil {
		return fmt.Errorf("couldn't fetch pokemonInfo for %s. Error: %w", name, err)
	}

	if _, ok := cfg.Pokedex[name]; !ok {
		fmt.Print("you have not caught that pokemon\n")
		return nil
	}

	fmt.Printf("Height: %d\n", pokemonInfo.Height)
	fmt.Printf("Weight: %d\n", pokemonInfo.Weight)
	fmt.Print("Stats:\n")
	for _, bstat := range pokemonInfo.Stats {
		fmt.Printf("  -%s: %d\n", bstat.Stat.Name, bstat.BaseStat)
	}
	fmt.Print("Types:\n")
	for _, pokeType := range pokemonInfo.Types {
		fmt.Printf("  - %s\n", pokeType.Type.Name)
	}

	return nil
}
