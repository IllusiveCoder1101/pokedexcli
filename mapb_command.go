package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandMapb(a *config, _non_usable string) error {

	var prev string

	if len(a.previous) == 0 {
		fmt.Println("you're on the first page")
		return nil
	} else {
		prev = a.previous
	}
	res, err := http.Get(prev)

	if err != nil {
		return fmt.Errorf("%v", err)
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return fmt.Errorf("%v", err)
	}

	data := pokemonData{}

	if err := json.Unmarshal(body, &data); err != nil {
		return fmt.Errorf("%v", err)
	}
	a.next = data.Next

	if len(data.Previous) == 0 {
		a.previous = data.Previous
	}

	for _, dat := range data.Results {
		fmt.Println(dat.Name)
	}

	return nil

}
