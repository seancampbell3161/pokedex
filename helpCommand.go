package main

import (
	"fmt"
)

func helpCommand(cfg *config, s *string, commands *map[string]cliCommand) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	for key, value := range *commands {
		fmt.Println(key + " - " + value.description)
	}

	return nil
}
