package main

import (
	"fmt"
)

func commandExplore(cfg *config, args []string) error {
	if args == nil || len(args) == 0 {
		fmt.Println("Which area do you want to explore?")
		return nil
	}

	locationsDetailsResp, err := cfg.pokeapiClient.LocationDetails(args[0])
	if err != nil {
		println("DEBUG: Error here!")
		return err
	}

	fmt.Println("Exploring ", locationsDetailsResp.Location.Name, "...")
	if len(locationsDetailsResp.PokemonEncounters) > 0 {
		fmt.Println("Found Pokemon:")
		for _, pokemon := range locationsDetailsResp.PokemonEncounters {
			fmt.Println(" - ", pokemon.Pokemon.Name)
		}
	} else {
		fmt.Println("No Pokemon found!")
	}

	return nil
}
