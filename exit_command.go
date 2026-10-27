package main

import (
	"fmt"
	"os"
)

func commandExit(a *config, _non_usable string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	defer os.Exit(0)
	return nil
}
