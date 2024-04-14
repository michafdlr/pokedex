package main

import (
	"fmt"
)

func ExploreLocation(cfg *config, name *string) error {
	pokeResp, err := cfg.PokeClient.GetPokemons(name)
	if err != nil {
		return err
	}
	for _, poke := range pokeResp.PokemonEncounters {
		fmt.Println(poke.Pokemon.Name)
	}
	return nil
}
