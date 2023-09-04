package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemonList(baseUrl string, areaName *string) (LocationArea, error) {
	value, ok := c.cache.GetCache(*areaName)
	location := LocationArea{}

	if !ok {
		res, err := http.Get(baseUrl + *areaName)
		if err != nil {
			return LocationArea{}, err
		}

		body, err := io.ReadAll(res.Body)
		json.Unmarshal(body, &location)

		c.cache.AddToCache(*areaName, body)

		return location, nil
	} else {
		fmt.Println("hit cache")
		json.Unmarshal(value, &location)
		return location, nil
	}
}
