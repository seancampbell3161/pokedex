package main

import "fmt"

func inspectCommand(cfg *config, pokemonName *string) error {
	pokemon, ok := cfg.pokedexEntries[*pokemonName]
	if !ok {
		fmt.Println("You have not caught this pokemon")
		return nil
	}

	fmt.Printf("Name: %v\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("\t-%v: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, pokeType := range pokemon.Types {
		fmt.Printf("\t-%v\n", pokeType.Type.Name)
	}

	return nil
}
