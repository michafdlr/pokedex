package main

import (
	"errors"
	"fmt"
)

func CommandMapfwd(cfg *config) error {
	locResp, err := cfg.PokeClient.GetLocations(cfg.Next)
	if err != nil {
		return err
	}
	cfg.Next = locResp.Next
	cfg.Previous = locResp.Previous
	for _, loc := range locResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func CommandMapbwd(cfg *config) error {
	if cfg.Previous == nil {
		return errors.New("you are on the first page")
	}
	locResp, err := cfg.PokeClient.GetLocations(cfg.Previous)
	if err != nil {
		return err
	}
	cfg.Next = locResp.Next
	cfg.Previous = locResp.Previous
	for _, loc := range locResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
