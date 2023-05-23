package main

import (
	"fmt"
	"pokeapis"
)

func displayPokemonList(cfg *config, area string) error {
	locations, err := pokeapis.ListPokemonLocations(area)
	if err != nil {
		fmt.Println(err)
		return err
	}
	cfg.locationArea = &locations.Name
	for _, value := range locations.Encounters {
		fmt.Println(value.Poke.PokeName)
	}
	return nil
}
