package main

import (
	"fmt"
	"maps"
	"slices"
	"github.com/alex-bartleynees/pokedex/internal/config"
)

func commandHelp(cfg *config.Config) error {
	names := slices.Sorted(maps.Keys(cfg.Commands))

	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, name := range names {
		fmt.Printf("%s: %s\n", name, cfg.Commands[name].Description)
	}
	return nil
}
