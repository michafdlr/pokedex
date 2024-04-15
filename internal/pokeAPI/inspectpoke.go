package pokeAPI

import "fmt"

func (c *Client) InspectPoke(pokename *string) (PokemonInfo, error) {
	if info, ok := c.pokedex[*pokename]; ok {
		return info, nil
	}
	return PokemonInfo{}, fmt.Errorf("%s not yet catched", *pokename)
}
