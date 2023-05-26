package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type config struct {
	nextUrl      *string
	prevUrl      *string
	locationArea *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
}

var cliCommands map[string]cliCommand

func startRepl(cfg *config, optional string) {
	cli := "pokedex >"
	var command string = ""
	cliCommands = updateCli()
	for {
		fmt.Printf("%s", cli)
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		command = scanner.Text()
		commands := strings.Split(command, " ")
		switch commands[0] {
		case "help":
			err := cliCommands["help"].callback(cfg, optional)
			if err != nil {
				fmt.Println(err)
				return
			}
		case "exit":
			err := cliCommands["exit"].callback(cfg, optional)
			if err != nil {
				fmt.Println(err)
				return
			}
			return
		case "map":
			err := cliCommands["map"].callback(cfg, optional)
			if err != nil {
				fmt.Println(err)
				return
			}
		case "mapb":
			err := cliCommands["mapb"].callback(cfg, optional)
			if err != nil {
				fmt.Println(err)
				return
			}
		case "explore":
			area := commands[1]
			err := cliCommands["explore"].callback(cfg, area)
			if err != nil {
				return
			}
		case "catch":
			pokemonName := commands[1]
			err := cliCommands["catch"].callback(cfg, pokemonName)
			if err != nil {
				return
			}
		default:
			fmt.Println("Please provide help or exit to find usage")
		}
	}
}

func updateCli() map[string]cliCommand {
	data := map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exits the poke dex cli",
			callback:    exitHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 area locations",
			callback:    displayLocations,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 locations",
			callback:    displayBackwardLocations,
		},
		"explore": {
			name:        "explore",
			description: "List of Pokemon in a given area",
			callback:    displayPokemonList,
		},
		"catch": {
			name:        "catch",
			description: "Catching some pokemon",
			callback:    catchPokemon,
		},
	}
	return data
}
