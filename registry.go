package main

import "pokedexcli/internal/pokecache"

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
}
type LocationArea struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type pokemonData struct {
	Count    int            `json:"count"`
	Next     string         `json:"next"`
	Previous string         `json:"previous"`
	Results  []LocationArea `json:"results"`
}
type PokemonDescription struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}
type PokemonEncounters struct {
	Pokemon PokemonDescription `json:"pokemon"`
}
type pokemonExploreData struct {
	Id                int                 `json:"id"`
	Name              string              `json:"name"`
	GameIndex         int                 `json:"game_index"`
	PokemonEncounters []PokemonEncounters `json:"pokemon_encounters"`
}
type StatDesc struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type PokemonStatsDesc struct {
	BaseStat int      `json:"base_stat"`
	Effort   int      `json:"effort"`
	Stat     StatDesc `json:"stat"`
}
type PokemonTypesDesc struct {
	Slot int      `json:"slot"`
	Type StatDesc `json:"type"`
}
type PokemonStats struct {
	Name           string             `json:"name"`
	Height         int                `json:"height"`
	Weight         int                `json:"weight"`
	BaseExperience int                `json:"base_experience"`
	Stats          []PokemonStatsDesc `json:"stats"`
	Types          []PokemonTypesDesc `json:"types"`
}
type config struct {
	commands map[string]cliCommand
	next     string
	previous string
	pokedex  map[string]PokemonStats
	cache    *pokecache.Cache
}
