package main

import (
	"fmt"
)

func mapCommand(config *config) error {
	locations, err := config.pokeapiClient.GetLocations(config.next)
	if err != nil {
		return err
	}

	config.next = locations.Next
	config.previous = locations.Previous

	for _, location := range locations.Results {
		fmt.Printf("%v\n", location.Name)
	}
	return nil
}

func mapbCommand(config *config) error {
	locations, err := config.pokeapiClient.GetLocations(config.previous)
	if err != nil {
		return err
	}

	config.next = locations.Next
	config.previous = locations.Previous

	for _, location := range locations.Results {
		fmt.Printf("%v\n", location.Name)
	}
	return nil
}
