package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/alex-bartleynees/pokedex/internal/config"
)

func startRepl(cfg *config.Config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		if !scanner.Scan() {
			break
		}

		input := cleanInput(scanner.Text())

		if len(input) == 0 {
			continue
		}

		command, exists := cfg.Commands[input[0]]
		if !exists {
			fmt.Println("Unknown command")
			continue
		}

		if err := command.Callback(cfg); err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(input string) []string {
	cleanedInput := strings.ToLower(input)
	cleanedInput = strings.TrimSpace(cleanedInput)
	return strings.Fields(cleanedInput)
}
