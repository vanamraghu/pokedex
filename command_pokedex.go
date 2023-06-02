package main

import (
	"fmt"
	"pokeapis"
)

func displayCaughtPokemon(cfg *config, optionalArg string) error {
	list, err := pokeapis.CaughtPokeMonList()
	if err != nil {
		return err
	}
	// Iterate through the list
	fmt.Println("Your Pokedex:")
	for _, val := range list {
		fmt.Printf(" - %s\n", val)
	}
	return nil
}
