package main

import (
	"fmt"
	"pokeapis"
)

func catchPokemon(cfg *config, optional string) error {
	exp, err := pokeapis.PokemonCatchDetails(optional)
	if err != nil {
		fmt.Println("Error is ", err)
		return err
	}
	fmt.Println(exp)
	return nil
}
