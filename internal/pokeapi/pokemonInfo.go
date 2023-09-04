package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemonInfo(name *string) (Pokemon, error) {
	baseUrl := "https://pokeapi.co/api/v2/pokemon/"
	finalUrl := baseUrl + *name

	pokemon := Pokemon{}
	response, err := http.Get(finalUrl)
	if err != nil {
		return Pokemon{}, nil
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return Pokemon{}, err
	}

	json.Unmarshal(body, &pokemon)

	return pokemon, nil
}
