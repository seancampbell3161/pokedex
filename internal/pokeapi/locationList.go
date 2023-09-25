package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocations(url *string) (LocationAreas, error) {
	finalUrl := "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"
	if url != nil {
		finalUrl = *url
	}

	locations := LocationAreas{}

	if val, ok := c.cache.GetCache(finalUrl); ok {
		err := json.Unmarshal(val, &locations)
		if err != nil {
			return locations, err
		}
		return locations, nil
	} else {
		res, err := http.Get(finalUrl)

		body, err := io.ReadAll(res.Body)
		err = res.Body.Close()
		if err != nil {
			return LocationAreas{}, err
		}

		locations := LocationAreas{}

		err = json.Unmarshal(body, &locations)
		if err != nil {
			return locations, err
		}

		c.cache.AddToCache(finalUrl, body)

		return locations, nil
	}

	return locations, nil
}
