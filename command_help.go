package main

import "fmt"

func commandHelp(cfg *config, name *string) error {
	helpText := "\nWelcome to pokedex\n\nUsage:\n\n"
	for _, cmd := range getCommands() {
		helpText += cmd.name + ": " + cmd.description + "\n\n"
	}
	fmt.Println(helpText)
	return nil
}
