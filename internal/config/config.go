package config

import (
	"github.com/alex-bartleynees/pokedex/internal/pokedexapi"
)

type CliCommand struct {
	Name        string
	Description string
	Callback    func(cfg *Config) error
}

type Config struct {
	Commands      map[string]CliCommand
	PokeApiClient *pokedexapi.Client
	NextPageURL    *string
	PreviousPageURL *string
}

func NewConfig(commands map[string]CliCommand, client *pokedexapi.Client) *Config {
	return &Config{Commands: commands, PokeApiClient: client}
}
