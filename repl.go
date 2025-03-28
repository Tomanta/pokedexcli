package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/tomanta/pokedexcli/internal/pokeAPI"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

type config struct {
	pokeapiClient        pokeAPI.Client
	pokedex              map[string]pokeAPI.PokemonDetails
	nextLocationsURL     *string
	previousLocationsURL *string
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
		"explore": {
			name:        "explore",
			description: "Explores a given location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Examine a pokemon you have caught",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List pokemon you have caught",
			callback:    commandPokedex,
		},		
	}
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	command_map := getCommands()

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())

		if len(words) == 0 {
			continue
		}

		cmd, exists := command_map[words[0]]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		if exists {
			err := cmd.callback(cfg, args...)
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
