package main

import (
	"github.com/seancampbell3161/pokedex/internal/pokeapi"
	"time"
)

func main() {
	pokeClient := pokeapi.NewClient(time.Second*5, time.Minute*5)
	cfg := &config{
		pokeapiClient:  pokeClient,
		pokedexEntries: make(map[string]pokeapi.Pokemon),
	}
	startRepl(cfg)
}
