package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/michafdlr/pokedexcli/internal/pokeAPI"
)

func main() {
	client := pokeAPI.NewClient(5*time.Second, 5*time.Minute)
	cfg := &config{
		PokeClient: client,
	}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}
		if scanner.Err() != nil {
			fmt.Println("Error occured while reading from stdin")
		}
		cmdName := words[0]
		cmd, exists := getCommands()[cmdName]
		if exists {
			err := cmd.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown Command")
			continue
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	PokeClient pokeAPI.Client
	Next       *string
	Previous   *string
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "get next 20 locations",
			callback:    CommandMapfwd,
		},
		"mapb": {
			name:        "mapb",
			description: "get previous 20 locations",
			callback:    CommandMapbwd,
		},
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
