package main

import (
	"bufio"
	"fmt"
	"github.com/seancampbell3161/pokedex/internal/pokeapi"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, *string) error
}

type config struct {
	pokeapiClient  pokeapi.Client
	next           *string
	previous       *string
	pokedexEntries map[string]pokeapi.Pokemon
}

func startRepl(cfg *config) {
	commands := getCommands()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		userInput := scanner.Text()
		splitWords := strings.Split(userInput, " ")

		if scanner.Text() == "exit" {
			return
		}
		if len(splitWords) > 1 {
			err := commands[splitWords[0]].callback(cfg, &splitWords[1])
			if err != nil {
				fmt.Println(err)
			}
		} else {
			err := commands[splitWords[0]].callback(cfg, &splitWords[0])

			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
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
		"explore": {
			name:        "explore",
			description: "Explore the area for pokemon",
			callback:    exploreCommand,
		},
		"catch": {
			name:        "catch",
			description: "attempt to catch a pokemon",
			callback:    catchCommand,
		},
	}
}
