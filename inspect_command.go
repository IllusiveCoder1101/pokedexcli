package main

import (
	"fmt"
)

func commandInspect(a *config, pokemonName string) error {
	pokedexData := a.pokedex

	value, ok := pokedexData[pokemonName]

	if !ok {
		fmt.Println("you have not caught that pokemon")
	} else {
		fmt.Printf("Name: %s\n", value.Name)
		fmt.Printf("Height: %d\n", value.Height)
		fmt.Printf("Weight: %d\n", value.Weight)
		fmt.Printf("Stats: \n")
		for i := 0; i < len(value.Stats); i++ {
			fmt.Printf(" -%s: %d\n", value.Stats[i].Stat.Name, value.Stats[i].BaseStat)
		}
		fmt.Printf("Types: \n")
		for i := 0; i < len(value.Types); i++ {
			fmt.Printf(" - %s\n", value.Types[i].Type.Name)
		}
	}

	return nil

}
