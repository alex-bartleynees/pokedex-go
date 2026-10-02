package main

import (
	"github.com/alex-bartleynees/pokedex/internal/config"
	"github.com/alex-bartleynees/pokedex/internal/pokedexapi"
)

func newConfig(client *pokedexapi.Client) *config.Config {
	return config.NewConfig(map[string]config.CliCommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex REPL",
			Callback:    commandExit,
		},
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    commandHelp,
		},
		"map": {
			Name:        "map",
			Description: "Displays the names of location areas in the Pokemon world",
			Callback:    commandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Go to the previous page of location areas in the Pokemon world",
			Callback:    commandMapBack,
		},
	}, client)
}
