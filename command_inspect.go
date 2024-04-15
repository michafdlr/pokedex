package main

import "fmt"

func CommandInspect(cfg *config, pokename *string) error {
	info, err := cfg.PokeClient.InspectPoke(pokename)
	if err != nil {
		return err
	}
	fmt.Println("Height:", info.Height)
	fmt.Println("Weight:", info.Weight)
	fmt.Println("Stats:")
	for _, val := range info.Stats {
		fmt.Println("  - "+val.Stat.Name+":", val.BaseStat)
	}
	fmt.Println("Types:")
	for _, val := range info.Types {
		fmt.Println("  - " + val.Type.Name)
	}
	return nil
}
