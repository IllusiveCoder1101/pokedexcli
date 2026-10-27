package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
)

func commandCatch(a *config, pokemonName string) error {

	res, err := http.Get(fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s/", pokemonName))

	if err != nil {
		return fmt.Errorf("%v", err)
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return fmt.Errorf("%v", err)
	}

	data := PokemonStats{}

	if err := json.Unmarshal(body, &data); err != nil {
		return fmt.Errorf("%v", err)
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	catchRate := float64(rand.Intn(data.BaseExperience)) / float64(data.BaseExperience)

	if catchRate >= .5 {
		fmt.Printf("%s was caught!\n", pokemonName)
		a.pokedex[pokemonName] = data
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}

	return nil

}
