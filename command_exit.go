package main

import (
	"fmt"
	"os"
	"github.com/alex-bartleynees/pokedex/internal/config"
)

func commandExit(cfg *config.Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
