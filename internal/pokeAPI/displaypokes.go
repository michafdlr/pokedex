package pokeAPI

import "fmt"

func (c *Client) GetPokedex() (map[string]PokemonInfo, error) {
	pokes := c.pokedex
	if len(c.pokedex) == 0 {
		return map[string]PokemonInfo{}, fmt.Errorf("no pokemons catched yet")
	}
	return pokes, nil
}
