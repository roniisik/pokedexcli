package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, args []string) error {
	if len(args) != 0 {
		return errors.New("pokedex command takes no args")
	}

	fmt.Println("Your Pokedex:")
	for key := range cfg.Pokedex {
		fmt.Printf(" - %s\n", key)
	}
	return nil
}
