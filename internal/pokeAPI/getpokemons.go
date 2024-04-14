package pokeAPI

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemons(name *string) (PokemonsByLocation, error) {
	url := baseURL + "/location-area" + "/1/" //default to first area
	if name != nil {
		url = baseURL + "/location-area/" + *name + "/"
	}
	if val, ok := c.cache.Get(url); ok {
		Pokemons := PokemonsByLocation{}
		err := json.Unmarshal(val, &Pokemons)
		if err != nil {
			return PokemonsByLocation{}, err
		}
		return Pokemons, nil
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonsByLocation{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonsByLocation{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonsByLocation{}, err
	}
	Pokemons := PokemonsByLocation{}
	err = json.Unmarshal(body, &Pokemons)
	if err != nil {
		return PokemonsByLocation{}, err
	}
	c.cache.Add(url, body)
	return Pokemons, nil
}
