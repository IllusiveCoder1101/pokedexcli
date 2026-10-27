package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandMap(a *config, _ string) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if len(a.next) != 0 {
		url = a.next
	}

	body, ok := a.cache.Get(url)
	if !ok {
		res, err := http.Get(url)
		if err != nil {
			return err
		}
		defer res.Body.Close()

		body, err = io.ReadAll(res.Body)
		if err != nil {
			return err
		}

		a.cache.Add(url, body)
	}

	data := pokemonData{}
	if err := json.Unmarshal(body, &data); err != nil {
		return err
	}

	a.next = data.Next
	a.previous = data.Previous

	for _, dat := range data.Results {
		fmt.Println(dat.Name)
	}

	return nil
}
