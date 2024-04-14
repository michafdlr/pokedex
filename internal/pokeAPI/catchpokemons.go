package pokeAPI

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
)

func (c *Client) CatchPokemon(name *string) (PokemonInfo, error) {
	if name == nil {
		return PokemonInfo{}, errors.New("no pokemon name provided")
	}
	url := baseURL + "/pokemon/" + *name
	if val, ok := c.cache.Get(url); ok {
		Pokemons := PokemonInfo{}
		err := json.Unmarshal(val, &Pokemons)
		if err != nil {
			return PokemonInfo{}, err
		}
		return Pokemons, nil
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonInfo{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonInfo{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonInfo{}, err
	}
	Pokemons := PokemonInfo{}
	randNum := rand.Intn(101)
	if randNum < 50 {
		return PokemonInfo{}, fmt.Errorf("could not catch %s", *name)
	}
	err = json.Unmarshal(body, &Pokemons)
	if err != nil {
		return PokemonInfo{}, err
	}
	c.cache.Add(url, body)
	return Pokemons, nil
}
