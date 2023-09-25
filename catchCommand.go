package main

import "fmt"
import "math/rand"

func catchCommand(cfg *config, pokemon *string, commands *map[string]cliCommand) error {
	fmt.Printf("throwing a pokeball at %v...\n", *pokemon)

	value, err := cfg.pokeapiClient.GetPokemonInfo(pokemon)
	if err != nil {
		return err
	}

	if value.ID == 0 {
		fmt.Println("Failed to catch")
	}

	successThreshold := 60
	attemptVal := rand.Intn(value.BaseExperience)

	if attemptVal > successThreshold {
		fmt.Printf("%v escaped!\n", value.Name)
		return nil
	}

	entry, ok := cfg.pokedexEntries[value.Name]
	if !ok {
		cfg.pokedexEntries[value.Name] = value
		fmt.Printf("%v was caught!\n", value.Name)
	} else {
		fmt.Printf("%v was caught!\n", entry.Name)
	}

	return nil
}
