package main

import (
	"testing"
	"time"

	"github.com/alex-bartleynees/pokedex/internal/config"
	"github.com/alex-bartleynees/pokedex/internal/pokedexapi"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "CHARMANDER",
			expected: []string{"charmander"},
		},
		{
			input:    "Pikachu  ",
			expected: []string{"pikachu"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			// error and continue to next case
			t.Errorf("cleanInput(%q) returned %d elements, expected %d", c.input, len(actual), len(c.expected))
			continue
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("cleanInput(%q)[%d] = %q, expected %q", c.input, i, word, expectedWord)
			}
		}
	}
}

func TestCommandRegistry(t *testing.T) {
	cases := []struct {
		input    string
		expected config.CliCommand
		exists   bool
	}{
		{
			input:    "exit",
			expected: config.CliCommand{Name: "exit", Description: "Exit the Pokedex REPL", Callback: commandExit},
			exists:   true,
		},
		{
			input:    "help",
			expected: config.CliCommand{Name: "help", Description: "Displays a help message", Callback: commandHelp},
			exists:   true,
		},
		{
			input:  "charmander",
			exists: false,
		},
	}

	cfg := newConfig(pokedexapi.NewClient(time.Second))

	for _, c := range cases {
		actual, exists := cfg.Commands[c.input]

		if exists != c.exists {
			t.Errorf("Commands[%q] existence = %v, expected %v", c.input, exists, c.exists)
		}

		if actual.Name != c.expected.Name || actual.Description != c.expected.Description {
			t.Errorf("Commands[%q] = %+v, expected %+v", c.input, actual, c.expected)
		}

		if actual.Callback == nil && c.expected.Callback != nil {
			t.Errorf("Commands[%q] callback is nil, expected non-nil", c.input)
		}
	}
}
