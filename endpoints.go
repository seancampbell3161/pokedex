package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func mapCommand(c *config) error {
	res, err := http.Get(c.next)

	if err != nil {
		return err
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	if err != nil {
		return err
	}

	locations := LocationAreas{}

	err = json.Unmarshal(body, &locations)

	c.next = locations.Next
	if locations.Previous != nil {
		c.previous = locations.Previous
	} else {
		c.previous = nil
	}

	if err != nil {
		return err
	}

	for _, location := range locations.Results {
		fmt.Printf("%v\n", location.Name)
	}

	return nil
}

func mapbCommand(c *config) error {
	if c.previous == nil {
		return errors.New("There are no previous areas")
	}

	res, err := http.Get(*c.previous)

	if err != nil {
		return err
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	if err != nil {
		return err
	}

	locations := LocationAreas{}

	err = json.Unmarshal(body, &locations)

	c.next = locations.Next

	if err != nil {
		return err
	}

	for _, location := range locations.Results {
		fmt.Printf("%v\n", location.Name)
	}

	return nil
}
