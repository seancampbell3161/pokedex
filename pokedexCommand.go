package main

import "fmt"

func pokedexCommand(cfg *config, s *string) error {
	fmt.Println("Your pokedex:")

	if len(cfg.pokedexEntries) == 0 {
		fmt.Println("<empty>")
		return nil
	}

	for k, _ := range cfg.pokedexEntries {
		fmt.Println(" - " + k)
	}

	return nil
}
