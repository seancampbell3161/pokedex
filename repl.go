package main

import (
	"bufio"
	"fmt"
	"github.com/seancampbell3161/pokedex/internal/pokeapi"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	pokeapiClient pokeapi.Client
	next          *string
	previous      *string
}

func startRepl(cfg *config) {
	commands := map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    helpCommand,
		},
		"exit": {
			name:        "exit",
			description: "Exits the pokedex",
			callback:    exitCommand,
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 locations at a time in the Pokemon world",
			callback:    mapCommand,
		},
		"mapb": {
			name:        "mapb",
			description: "Goes back to the previous locations names",
			callback:    mapbCommand,
		},
	}
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		for scanner.Scan() {
			switch st := scanner.Text(); st {
			case "help":
				err := commands["help"].callback(cfg)
				if err != nil {
					fmt.Printf("%v\n", err)
				}
			case "exit":
				err := commands["exit"].callback(cfg)
				if err != nil {
					fmt.Printf("%v\n", err)
				}
				return
			case "map":
				err := commands["map"].callback(cfg)
				if err != nil {
					fmt.Printf("%v\n", err)
				}
			case "mapb":
				err := commands["mapb"].callback(cfg)
				if err != nil {
					fmt.Printf("%v\n\n", err)
				}
			}
			fmt.Print("Pokedex > ")
		}
	}
}
