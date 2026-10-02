package main 

import (
	"github.com/alex-bartleynees/pokedex/internal/config"
	"github.com/alex-bartleynees/pokedex/internal/pokedexapi"
	"fmt"
)

func commandMap(cfg *config.Config) error {
	return showLocationPage(cfg, cfg.NextPageURL)
}

func commandMapBack(cfg *config.Config) error {
	if cfg.PreviousPageURL == nil {
		fmt.Println("No previous page available.")
		return nil
	}
	
	return showLocationPage(cfg, cfg.PreviousPageURL)

}

func showLocationPage(cfg *config.Config, url *string) error {
	var page, err = cfg.PokeApiClient.GetLocationPage(url)

	if err != nil {
		return err
	}

	printResults(page)

	setPageURLs(cfg, page)

	return nil
}

func printResults(page *pokedexapi.LocationPage) {
	for _, area := range page.Results {
		fmt.Println(area.Name)
	}
}

func setPageURLs(cfg *config.Config, page *pokedexapi.LocationPage) {
	cfg.NextPageURL = page.Next
	cfg.PreviousPageURL = page.Previous
}

