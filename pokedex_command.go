package main

import (
	"fmt"
)

func commandPokedex(a *config, _non_usable string) error {
	pokedexData := a.pokedex
	fmt.Println("Your Pokedex: ")
	for _, value := range pokedexData {
		fmt.Printf(" - %s\n", value.Name)
	}

	return nil

}
