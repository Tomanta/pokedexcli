package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	fmt.Println("NOT YET COMPLETED")
	if len(args) != 1 {
		fmt.Println("Which pokemon do you want to catch?")
		return nil
	}

	pokemonDetailsResp, err := cfg.pokeapiClient.PokemonDetails(args[0])
	if err != nil {
		// TODO: If lookup fails, the error appears here
		// "invalid character 'N' looking for beginning of value"
		return err
	}

	pokemon := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)

	// TODO: Fetch Pokemon data from API
	chance := 0.9 - (float64(pokemonDetailsResp.BaseExperience) / 1000) // TODO: Modify chance based on Pokemon's base experience
	fmt.Printf("%.02f chance to catch...", chance)
	if rand.Float64() > chance {
		fmt.Printf("%s escaped!\n", pokemon)
	} else {
		fmt.Printf("%s was caught!\n", pokemon)
		// TODO: Add to pokedex
	}
	return nil
}
