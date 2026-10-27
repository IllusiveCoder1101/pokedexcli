package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandExplore(a *config, regionName string) error {

	res, err := http.Get(fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s/", regionName))

	if err != nil {
		return fmt.Errorf("%v", err)
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return fmt.Errorf("%v", err)
	}

	data := pokemonExploreData{}

	if err := json.Unmarshal(body, &data); err != nil {
		return fmt.Errorf("%v", err)
	}

	fmt.Printf("Exploring %s...\n", regionName)
	fmt.Println("Found Pokemon:")
	for _, dat := range data.PokemonEncounters {
		fmt.Printf("  - %s\n", dat.Pokemon.Name)
	}

	return nil

}
