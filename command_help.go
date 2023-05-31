package main

import "fmt"

func commandHelp(cfg *config, optional string) error {
	fmt.Println("Welcome to the pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	fmt.Printf("help:%v\n", cliCommands["help"].description)
	fmt.Printf("exit:%v\n", cliCommands["exit"].description)
	fmt.Printf("map:%s\n", cliCommands["map"].description)
	fmt.Printf("mapb:%s\n", cliCommands["mapb"].description)
	fmt.Printf("explore <area-name>:%v\n", cliCommands["explore"].description)
	fmt.Printf("catch <pokemon-name>: %s\n", cliCommands["catch"].description)
	fmt.Printf("inspect <pokemon-name>: %s\n", cliCommands["inspect"].description)
	fmt.Println("")
	return nil
}
