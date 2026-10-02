package main

import (
	"fmt"
	"github.com/alex-bartleynees/pokedex/internal/config"
	"os"
)

func commandExit(cfg *config.Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
