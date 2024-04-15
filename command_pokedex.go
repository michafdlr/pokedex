package main

import "fmt"

func CommandPokedex(cfg *config, name *string) error {
	pokemons, err := cfg.PokeClient.GetPokedex()
	if err != nil {
		return err
	}
	fmt.Println("Your Pokedex:")
	for _, poke := range pokemons {
		fmt.Println("  - " + poke.Name)
	}
	return nil
}
