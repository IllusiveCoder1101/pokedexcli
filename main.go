package main

import (
	"pokedexcli/internal/pokecache"
	"time"
)

func main() {

	cache := pokecache.NewCache(5 * time.Second)

	cfg := &config{
		commands: make(map[string]cliCommand),
		pokedex:  make(map[string]PokemonStats),
		cache:    &cache,
	}

	// Register commands into the config state
	cfg.commands["exit"] = cliCommand{
		name:        "exit",
		description: "Exit the REPL",
		callback:    commandExit, // 5. Updated callback
	}
	cfg.commands["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp, // 5. Updated callback
	}

	cfg.commands["map"] = cliCommand{
		name:        "map",
		description: "Get 20 location areas in pokemon world",
		callback:    commandMap, // 5. Updated callback
	}

	cfg.commands["mapb"] = cliCommand{
		name:        "mapb",
		description: "Gets back the old 20 location areas in pokemon world",
		callback:    commandMapb, // 5. Updated callback
	}

	cfg.commands["explore"] = cliCommand{
		name:        "explore",
		description: "Explore location areas in pokemon world and get pokemon available in that area",
		callback:    commandExplore, // 5. Updated callback
	}

	cfg.commands["catch"] = cliCommand{
		name:        "catch",
		description: "Catch pokemon to save in pokedex",
		callback:    commandCatch, // 5. Updated callback
	}

	cfg.commands["inspect"] = cliCommand{
		name:        "inspect",
		description: "Get pokemon details saved in pokedex",
		callback:    commandInspect, // 5. Updated callback
	}

	cfg.commands["pokedex"] = cliCommand{
		name:        "pokedex",
		description: "List all the pokemons captured",
		callback:    commandPokedex, // 5. Updated callback
	}

	repl(cfg)
}
