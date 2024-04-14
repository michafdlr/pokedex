package main

import (
	"fmt"
)

func CatchPokemon(cfg *config, name *string) error {
	fmt.Printf("Throwing ball at %s\n\n", *name)
	_, err := cfg.PokeClient.CatchPokemon(name)
	if err != nil {
		return err
	}
	fmt.Printf("\nCatched %s\n\n", *name)
	return nil
}
