package main

import (
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		fmt.Println("Which area do you want to explore?")
		return nil
	}

	locationsDetailsResp, err := cfg.pokeapiClient.LocationDetails(args[0])
	if err != nil {
		// TODO: If lookup fails, the error appears here
		// "invalid character 'N' looking for beginning of value"
		return err
	}

	fmt.Printf("Exploring %s...\n", locationsDetailsResp.Location.Name)
	if len(locationsDetailsResp.PokemonEncounters) > 0 {
		fmt.Println("Found Pokemon:")
		for _, pokemon := range locationsDetailsResp.PokemonEncounters {
			fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
		}
	} else {
		fmt.Println("No Pokemon found!")
	}

	return nil
}
