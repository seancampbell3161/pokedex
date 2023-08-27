package main

import (
	"fmt"
)

func helpCommand(*config) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")
	fmt.Println("map: Displays a list of 20 locations in Pokemon, repeated entries will retrieve the next 20 locations")
	fmt.Print("mapb: Goes back to the previous 20 locations listed")
	fmt.Println()

	return nil
}
