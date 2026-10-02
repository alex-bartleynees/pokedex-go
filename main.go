package main

import (
	"time"

	"github.com/alex-bartleynees/pokedex/internal/pokedexapi"
)

func main() {
	client := pokedexapi.NewClient(10 * time.Second)
	cfg := newConfig(client)
	startRepl(cfg)
}
