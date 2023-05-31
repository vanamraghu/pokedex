package main

import (
	"pokeapis"
)

func displayInspectStats(cfg *config, pokeName string) error {
	_, err := pokeapis.ListPokemonStats(pokeName)
	if err != nil {
		return err
	}
	//fmt.Println(string(stats))
	return nil
}
