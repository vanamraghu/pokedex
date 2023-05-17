package main

import "fmt"

type config struct {
	nextUrl *string
	prevUrl *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

var cliCommands map[string]cliCommand

func startRepl(cfg *config) {
	cli := "pokedex >"
	var command string = ""
	cliCommands = updateCli()
	for {
		fmt.Printf("%s", cli)
		_, _ = fmt.Scanln(&command)
		switch command {
		case "help":
			err := cliCommands["help"].callback(cfg)
			if err != nil {
				fmt.Println(err)
				return
			}
		case "exit":
			err := cliCommands["exit"].callback(cfg)
			if err != nil {
				fmt.Println(err)
				return
			}
			return
		case "map":
			err := cliCommands["map"].callback(cfg)
			if err != nil {
				fmt.Println(err)
				return
			}
		case "mapb":
			err := cliCommands["mapb"].callback(cfg)
			if err != nil {
				fmt.Println(err)
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
	}
	return data
}
