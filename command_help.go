package main

import "fmt"

func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	fmt.Printf("help:%v\n", cliCommands["help"].description)
	fmt.Printf("exit:%v\n", cliCommands["exit"].description)
	fmt.Printf("map:%s\n", cliCommands["map"].description)
	fmt.Printf("mapb:%s\n", cliCommands["mapb"].description)
	fmt.Println("")
	return nil
}
