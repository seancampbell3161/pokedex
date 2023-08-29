package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemonList(baseUrl string, areaName *string) (LocationArea, error) {
	location := LocationArea{}
	res, err := http.Get(baseUrl + *areaName)
	if err != nil {
		return LocationArea{}, err
	}

	body, err := io.ReadAll(res.Body)
	json.Unmarshal(body, &location)

	return location, nil
}
