package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func (c cliCommand) executeCommand() {

}

func startRepl() {
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
				err := commands["help"].callback()
				if err != nil {
					fmt.Printf("%v\n", err)
				}
			case "exit":
				err := commands["exit"].callback()
				if err != nil {
					fmt.Printf("%v\n", err)
				}
				return
			}
			fmt.Print("Pokedex > ")
		}
	}
}
