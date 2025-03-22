package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config) error
}

type config struct {
	next     string
	previous string
}

func getCommands() map[string]cliCommand {

	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays 20 location areas in the Pokemon world. Subsequent calls display the next 20 locations.",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 locations in the Pokemon world.",
			callback:    commandMapb,
		},
	}
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	command_map := getCommands()

	cfg := config{next: "", previous: ""}

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())

		if len(words) == 0 {
			continue
		}

		cmd, exists := command_map[words[0]]
		if exists {
			err := cmd.callback(&cfg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}

	}

}

func cleanInput(text string) []string {
	result := []string{}
	split_string := strings.Fields(text)
	for i := range split_string {
		result = append(result, strings.ToLower(split_string[i]))
	}

	return result
}
