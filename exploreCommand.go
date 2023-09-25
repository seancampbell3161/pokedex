package main

import (
	"fmt"
)

func exploreCommand(cfg *config, areaName *string, commands *map[string]cliCommand) error {
	baseUrl := "https://pokeapi.co/api/v2/location-area/"

	locationArea, err := cfg.pokeapiClient.GetPokemonList(baseUrl, areaName)
	if err != nil {
		return err
	}

	for _, pokemons := range locationArea.PokemonEncounters {
		fmt.Println(pokemons.Pokemon.Name)
	}

	return nil
}
