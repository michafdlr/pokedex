package main

import "fmt"

func CommandInspect(cfg *config, pokename *string) error {
	info, err := cfg.PokeClient.InspectPoke(pokename)
	if err != nil {
		return err
	}
	fmt.Println("- height:", info.Height)
	fmt.Println("- weight:", info.Weight)
	fmt.Println("- Stats:")
	for _, val := range info.Stats {
		fmt.Println("  - "+val.Stat.Name+":", val.BaseStat)
	}
	return nil
}
